package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type CountVector struct {
	Lin int
	Col int
	Id  string
}

func NewCountVector(lin int, col int, val string) CountVector {
	exp := CountVector{lin, col, val}
	return exp
}

func (p CountVector) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
