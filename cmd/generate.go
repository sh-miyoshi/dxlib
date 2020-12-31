package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type dxFunc struct {
	Name     string
	Response string
	GoArgs   string
	PArgs    string
}

type argument struct {
	Name   string
	GoType string
}

type comment struct {
	funcName string
	comments []string
}

var (
	cTypes = []string{"arrayint", "arraychar", "int *", "int", "unsigned int", "char *", "char", "double", "float", "LONGLONG"}

	// Sort in the same order as cTypes
	goTypes = []string{"[]int32", "[]byte", "*int32", "int32", "uint32", "string", "byte", "float64", "float32", "int64"}

	comments = []comment{}
)

func convToGoType(cType string) (string, bool) {
	for i, t := range cTypes {
		if t == cType {
			return goTypes[i], true
		}
	}
	return "", false
}

func parseArg(argStr string) ([]argument, error) {
	res := []argument{}
	if argStr == "" {
		return res, nil
	}

	for _, arg := range strings.Split(argStr, ",") {
		arg = strings.TrimSpace(arg)
		if arg == "..." {
			return nil, errors.New("... is unsupported")
		}

		ok := false
		for _, typ := range cTypes {
			if strings.HasPrefix(arg, typ) {
				ok = true
				name := strings.TrimPrefix(arg, typ)
				typ, _ = convToGoType(typ)
				res = append(res, argument{Name: name, GoType: typ})
				break
			}
		}
		if !ok {
			return nil, fmt.Errorf("Failed to find valid c type in arg %s", arg)
		}
	}

	return res, nil
}

func parse(line string) (*dxFunc, error) {
	if !strings.HasPrefix(line, "//dxlib") {
		return nil, nil
	}

	res := dxFunc{}
	raw := strings.TrimPrefix(line, "//dxlib ")

	// valid format is c format function
	//dxlib <response> <FuncName>(<arg1> interface{}, <arg2> ...)
	v1 := strings.Split(raw, " ")
	if len(v1) < 2 {
		return nil, fmt.Errorf("Failed to get response and function")
	}
	res.Response = v1[0]
	index := 1
	if v1[0] == "unsigned" {
		res.Response += " " + v1[1]
		index++
	}
	var ok bool
	res.Response, ok = convToGoType(res.Response)
	if !ok {
		return nil, fmt.Errorf("Invalid response type")
	}

	raw = ""
	for i := index; i < len(v1); i++ {
		raw += v1[i] + " "
	}

	v2 := strings.Split(raw, "(")
	if len(v2) != 2 {
		return nil, fmt.Errorf("Failed to find func start (")
	}
	res.Name = v2[0]

	argStr := strings.TrimSuffix(strings.TrimSpace(v2[1]), ")")
	args, err := parseArg(argStr)
	if err != nil {
		return nil, err
	}
	for _, arg := range args {
		res.GoArgs += arg.Name + " " + arg.GoType + ", "
		if strings.Contains(arg.GoType, "[]") {
			res.PArgs += "parray" + strings.Trim(arg.GoType, "[]") + "(" + arg.Name + "), "
		} else if strings.Contains(arg.GoType, "*") {
			res.PArgs += "pp" + strings.Trim(arg.GoType, "*") + "(" + arg.Name + "), "
		} else {
			res.PArgs += "p" + arg.GoType + "(" + arg.Name + "), "
		}
	}
	res.GoArgs = strings.TrimSuffix(res.GoArgs, ", ")
	res.PArgs = strings.TrimSuffix(res.PArgs, ", ")

	return &res, nil
}

func parseExtraFunc(line string) *dxFunc {
	if !strings.HasPrefix(line, "//ext_dxlib") {
		return nil
	}

	// valid format is only function name
	//ext_dxlib <FuncName>

	res := &dxFunc{
		Name: strings.TrimPrefix(line, "//ext_dxlib "),
	}
	return res
}

func parseComment(line string) error {
	if !strings.HasPrefix(line, "//comment") {
		return nil
	}

	// valid format
	//comment; <FuncName>; <comments>
	data := strings.Split(line, ";")
	if len(data) != 3 {
		return fmt.Errorf("invalid comment format")
	}

	comments = append(comments, comment{
		funcName: strings.TrimSpace(data[1]),
		comments: strings.Split(strings.TrimSpace(data[2]), "\\n"),
	})

	return nil
}

func getComment(funcName string) string {
	for _, c := range comments {
		if c.funcName == funcName {
			res := "// " + funcName + " "
			for _, msg := range c.comments {
				res += msg + "\n// "
			}
			return strings.TrimSuffix(res, "// ")
		}
	}
	return ""
}

func generate(w io.Writer, packageName string, funcs, extra []dxFunc) {
	procs := []string{}
	for _, f := range funcs {
		procs = append(procs, f.Name)
	}
	for _, f := range extra {
		procs = append(procs, f.Name)
	}

	data := map[string]interface{}{
		"Package":   packageName,
		"Procs":     procs,
		"Functions": funcs,
	}

	funcMap := map[string]interface{}{
		"getComment": getComment,
	}

	t := template.Must(template.New("main").Funcs(funcMap).Parse(outTemplate))
	t.Execute(w, data)
}

func generateExtFunc(w io.Writer) {
	io.WriteString(w, `
func DrawFormatString(x int32, y int32, color uint32, format string, a ...interface{}) int32 {
	str := fmt.Sprintf(format, a...)
	return DrawString(x, y, str, color, 0)
}

func DrawFormatStringToHandle(x int32, y int32, color uint32, fontHandle int32, format string, a ...interface{}) int32 {
	str := fmt.Sprintf(format, a...)
	return DrawStringToHandle(x, y, str, color, fontHandle, 0, FALSE)
}
`)
}

func main() {
	var input, output string
	flag.StringVar(&input, "input", "", "input file name")
	flag.StringVar(&input, "i", "", "input file name")
	flag.StringVar(&output, "output", "", "output file name")
	flag.StringVar(&output, "o", "", "output file name")

	flag.Parse()

	if input == "" || output == "" {
		fmt.Printf("Please set input and output\n")
		os.Exit(1)
	}

	fp, err := os.Open(input)
	if err != nil {
		fmt.Printf("Failed to open file %s: %v\n", input, err)
		os.Exit(1)
	}
	defer fp.Close()

	var data, extra []dxFunc
	packageName := ""
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "package") {
			packageName = strings.Split(line, " ")[1]
			continue
		}

		// Parse function
		d, err := parse(line)
		if err != nil {
			fmt.Printf("Failed to parse line %s: %v\n", line, err)
			os.Exit(1)
		}
		if d != nil {
			data = append(data, *d)
			continue
		}

		// Parse  extra function
		if d := parseExtraFunc(line); d != nil {
			extra = append(extra, *d)
			continue
		}

		// Parse command
		if err := parseComment(line); err != nil {
			fmt.Printf("Failed to parse line %s: %v\n", line, err)
			os.Exit(1)
		}
	}

	var buf bytes.Buffer
	generate(&buf, packageName, data, extra)
	generateExtFunc(&buf)
	ioutil.WriteFile(output, buf.Bytes(), 0644)
}

const outTemplate = `
// Code generated by 'go generate'; DO NOT EDIT.

package {{ .Package }}

import (
	"fmt"
	"syscall"
	"unsafe"
	"math"

	"golang.org/x/text/encoding/japanese"
    "golang.org/x/text/transform"
)

var (
	{{ range .Procs -}}
	dx_{{.}} *syscall.LazyProc
	{{ end }}
)

func Init(dllFile string) {
	mod := syscall.NewLazyDLL(dllFile)

	{{ range .Procs -}}
	dx_{{.}} = mod.NewProc("dx_{{.}}")
	{{ end }}
}

{{ range $i, $func := .Functions }}
{{ getComment $func.Name -}}
func {{ $func.Name }}({{ $func.GoArgs }}) {{ $func.Response }} {
	if dx_{{ $func.Name }} == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_{{ $func.Name }}.Call({{ $func.PArgs }})
	return  {{ $func.Response }}(res)
}
{{ end }}

func ppint32(i *int32) uintptr {
	return uintptr(unsafe.Pointer(i))
}

func pint32(i int32) uintptr {
	return uintptr(i)
}

func puint32(ui uint32) uintptr {
	return uintptr(ui)
}

func pstring(str string) uintptr {
	sjisStr, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), str)
	if err != nil {
		panic(err)
	}
	pbyte, err := syscall.BytePtrFromString(sjisStr)
	if err != nil {
		panic(err)
	}
	return uintptr(unsafe.Pointer(pbyte))
}

func pfloat32(f float32) uintptr {
	return uintptr(math.Float32bits(f))
}

func pfloat64(f float64) uintptr {
	return uintptr(math.Float64bits(f))
}

func pint64(i int64) uintptr {
	return uintptr(i)
}

func parraybyte(b []byte) uintptr {
	return uintptr(unsafe.Pointer(&b[0]))
}

func parrayint32(i []int32) uintptr {
	return uintptr(unsafe.Pointer(&i[0]))
}
`
