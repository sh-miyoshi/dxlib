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

// Data ...
type Data struct {
	PackageName string

	Funcs    []DxFunc
	ExtFuncs []DxFunc
}

// Parser ...
type Parser struct {
	fp *os.File
}

type argument struct {
	Name   string
	GoType string
}

var (
	cTypes = []string{"arrayint", "arraychar", "int *", "int", "unsigned int", "char *", "char", "double", "float", "LONGLONG"}
	// Sort in the same order as cTypes
	goTypes = []string{"[]int32", "[]byte", "*int32", "int32", "uint32", "string", "byte", "float64", "float32", "int64"}
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
			d, err := parseFunc(line)
			if err != nil {
				return nil, err
			}
			res.Funcs = append(res.Funcs, *d)
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

func parseFunc(line string) (*DxFunc, error) {
	res := DxFunc{}
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

func convertComment(funcName, comment string) string {
	res := "// " + funcName + " "
	for _, msg := range strings.Split(comment, "\\n") {
		res += msg + "\n// "
	}
	return strings.TrimSuffix(res, "// ")
}
