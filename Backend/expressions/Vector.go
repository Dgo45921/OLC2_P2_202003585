package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type Vector struct {
	Lin   int
	Col   int
	Value []interface{}
}

func NewVector(lin int, col int, val []interface{}) Vector {
	exp := Vector{lin, col, val}
	return exp
}

func (p Vector) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
