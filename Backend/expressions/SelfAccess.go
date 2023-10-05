package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type SelfAccess struct {
	Lin int
	Col int
	ID  string
}

func NewSelfAccess(lin int, col int, id string) SelfAccess {
	structaccess := SelfAccess{lin, col, id}
	return structaccess
}

func (p SelfAccess) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
