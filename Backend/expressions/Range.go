package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type Range struct {
	Lin        int
	Col        int
	FirstIndex interfaces.Expression
	LastIndex  interfaces.Expression
}

func NewRange(lin int, col int, findex interfaces.Expression, lindex interfaces.Expression) Range {
	exp := Range{lin, col, findex, lindex}
	return exp
}

func (p Range) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}