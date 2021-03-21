package parse

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// DxFunc ...
type DxFunc struct {
	Name     string
	Response string
	GoArgs   string
	PArgs    string
	Comment  string
}

// Argument ...
type Argument struct {
	Name    string
	GoType  string
	Default string
}

// Option ...
type Option struct {
	FuncName string
	Args     []Argument
}

// Data ...
type Data struct {
	PackageName string

	Funcs    []DxFunc
	ExtFuncs []DxFunc
	Options  []Option
}

// Parser ...
type Parser struct {
	fp *os.File
}

var (
	cTypes = []string{"arrayint", "arraychar", "int *", "int", "unsigned int", "char *", "char", "double", "float", "LONGLONG", "VECTOR", "HANDLE"}
	// Sort in the same order as cTypes
	goTypes = []string{"[]int", "[]byte", "*int", "int", "uint", "string", "byte", "float64", "float32", "int64", "VECTOR", "*int"}
)

func convToGoType(cType string) (string, bool) {
	for i, t := range cTypes {
		if t == cType {
			return goTypes[i], true
		}
	}
	return "", false
}

// New ...
func New(filePath string) (*Parser, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open input file %w", err)
	}

	return &Parser{
		fp: fp,
	}, nil
}

// End ...
func (p *Parser) End() {
	if p.fp != nil {
		p.fp.Close()
	}
}

// Parse ...
func (p *Parser) Parse() (*Data, error) {
	res := Data{}
	comments := make(map[string]string)

	scanner := bufio.NewScanner(p.fp)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "package") {
			res.PackageName = strings.Split(line, " ")[1]
			continue
		}
		if strings.HasPrefix(line, "//dxlib") {
			d, opt, err := parseFunc(line)
			if err != nil {
				return nil, err
			}
			res.Funcs = append(res.Funcs, *d)
			res.Options = append(res.Options, *opt)
			continue
		}
		if strings.HasPrefix(line, "//ext_dxlib") {
			d := parseExtraFunc(line)
			res.ExtFuncs = append(res.ExtFuncs, *d)
			continue
		}
		if strings.HasPrefix(line, "//comment") {
			funcName, comment, err := parseComment(line)
			if err != nil {
				return nil, err
			}
			comments[funcName] = comment
		}
	}

	// set comments
COMMENT:
	for funcName, comment := range comments {
		for i, f := range res.Funcs {
			if f.Name == funcName {
				res.Funcs[i].Comment = comment
				continue COMMENT
			}
		}
		for i, ef := range res.ExtFuncs {
			if ef.Name == funcName {
				res.ExtFuncs[i].Comment = comment
				continue COMMENT
			}
		}
		return nil, fmt.Errorf("No such function %s for comment", funcName)
	}

	return &res, nil
}

func parseFunc(line string) (*DxFunc, *Option, error) {
	res := DxFunc{}
	raw := strings.TrimPrefix(line, "//dxlib ")

	// valid format is c format function
	//dxlib <response> <FuncName>(<arg1> interface{}, <arg2> ...)
	v1 := strings.Split(raw, " ")
	if len(v1) < 2 {
		return nil, nil, fmt.Errorf("Failed to get response and function")
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
		return nil, nil, fmt.Errorf("Invalid response type")
	}

	raw = ""
	for i := index; i < len(v1); i++ {
		raw += v1[i] + " "
	}

	v2 := strings.Split(raw, "(")
	if len(v2) != 2 {
		return nil, nil, fmt.Errorf("Failed to find func start (")
	}
	res.Name = v2[0]

	argStr := strings.TrimSuffix(strings.TrimSpace(v2[1]), ")")
	args, err := parseArg(argStr)
	if err != nil {
		return nil, nil, err
	}

	opt := Option{}
	for _, arg := range args {
		if strings.Contains(arg.GoType, "[]") {
			res.PArgs += "parray" + strings.Trim(arg.GoType, "[]") + "(" + arg.Name + "), "
		} else if strings.Contains(arg.GoType, "*") {
			res.PArgs += "pp" + strings.Trim(arg.GoType, "*") + "(" + arg.Name + "), "
		} else {
			res.PArgs += "p" + arg.GoType + "(" + arg.Name + "), "
		}

		if arg.Default != "" {
			opt.Args = append(opt.Args, arg)
			continue
		}

		res.GoArgs += arg.Name + " " + arg.GoType + ", "
	}
	if len(opt.Args) > 0 {
		opt.FuncName = res.Name
		res.GoArgs += "opt ..." + opt.FuncName + "Option"
	}

	res.GoArgs = strings.TrimSuffix(res.GoArgs, ", ")
	res.PArgs = strings.TrimSuffix(res.PArgs, ", ")

	return &res, &opt, nil
}

func parseExtraFunc(line string) *DxFunc {
	// valid format is only function name
	//ext_dxlib <FuncName>

	res := &DxFunc{
		Name: strings.TrimPrefix(line, "//ext_dxlib "),
	}
	return res
}

func parseComment(line string) (string, string, error) {
	// valid format
	//comment; <FuncName>; <comments>
	data := strings.Split(line, ";")
	if len(data) != 3 {
		return "", "", fmt.Errorf("invalid comment format")
	}

	funcName := strings.TrimSpace(data[1])
	comment := convertComment(funcName, strings.TrimSpace(data[2]))

	return funcName, comment, nil
}

func parseArg(argStr string) ([]Argument, error) {
	res := []Argument{}
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
				val := ""
				if strings.Contains(name, "=") {
					// The argument has default value
					t := strings.Split(name, "=")
					name = strings.TrimSpace(t[0])
					val = strings.TrimSpace(t[1])
				}

				typ, _ = convToGoType(typ)
				res = append(res, Argument{Name: name, GoType: typ, Default: val})
				break
			}
		}
		if !ok {
			return nil, fmt.Errorf("Failed to find valid c type in arg %s", arg)
		}
	}

	return res, nil
}

func convertComment(funcName, comment string) string {
	res := "// " + funcName + " "
	for _, msg := range strings.Split(comment, "\\n") {
		res += msg + "\n// "
	}
	return strings.TrimSuffix(res, "// ")
}
