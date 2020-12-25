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

var (
	cTypes = []string{"arrayint", "arraychar", "int *", "int", "unsigned int", "char *", "char", "double", "float", "LONGLONG"}

	// Sort in the same order as cTypes
	goTypes = []string{"[]int", "[]byte", "*int", "int", "uint", "string", "byte", "float64", "float32", "int64"}
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
			return nil, errors.New("TODO ... is not implemented yet")
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

func generate(w io.Writer, packageName string, funcs []dxFunc) {
	procs := []string{}
	for _, f := range funcs {
		procs = append(procs, f.Name)
	}

	data := map[string]interface{}{
		"Package":   packageName,
		"Procs":     procs,
		"Functions": funcs,
	}

	t := template.Must(template.New("main").Parse(outTemplate))
	t.Execute(w, data)
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

	data := []dxFunc{}
	packageName := ""
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "package") {
			packageName = strings.Split(line, " ")[1]
			continue
		}
		d, err := parse(line)
		if err != nil {
			fmt.Printf("Failed to parse line %s: %v\n", line, err)
			os.Exit(1)
		}
		if d != nil {
			data = append(data, *d)
		}
	}

	var buf bytes.Buffer
	generate(&buf, packageName, data)
	ioutil.WriteFile(output, buf.Bytes(), 0644)
}

const outTemplate = `
// Code generated by 'go generate'; DO NOT EDIT.

package {{ .Package }}

import (
	"syscall"
	"unsafe"

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
func {{ $func.Name }}({{ $func.GoArgs }}) {{ $func.Response }} {
	if dx_{{ $func.Name }} == nil {
		panic("Please call dxlib.Init() at first")
	}

	res, _, _ := dx_{{ $func.Name }}.Call({{ $func.PArgs }})
	return  {{ $func.Response }}(res)
}
{{ end }}

func ppint(i *int) uintptr {
	return uintptr(unsafe.Pointer(i))
}

func pint(i int) uintptr {
	return uintptr(i)
}

func puint(ui uint) uintptr {
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
	return uintptr(f)
}

func pfloat64(f float64) uintptr {
	return uintptr(f)
}

func pint64(i int64) uintptr {
	return uintptr(i)
}

func parraybyte(b []byte) uintptr {
	return uintptr(unsafe.Pointer(&b[0]))
}

func parrayint(i []int) uintptr {
	return uintptr(unsafe.Pointer(&i[0]))
}
`
