package object

import (
	"bytes"
	"strings"

	"github.com/AlyxPink/meowlang/ast"
)

const FUNCTION_OBJ = "FUNCTION"

type Identifier struct {
	Name string
}

type Function struct {
	Parameters []*Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.Name)
	}

	out.WriteString("meow")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.TokenLiteral())
	out.WriteString("\n}")

	return out.String()
}
