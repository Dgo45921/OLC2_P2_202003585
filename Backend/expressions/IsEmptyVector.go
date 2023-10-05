package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type IsEmptyVector struct {
	Lin int
	Col int
	Id  string
}

func NewIsEmptyVector(lin int, col int, val string) IsEmptyVector {
	exp := IsEmptyVector{lin, col, val}
	return exp
}

func (p IsEmptyVector) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
