package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type UnarySum struct {
	Lin        int
	Col        int
	ID         string
	Op         string
	Expression interfaces.Expression
}

func NewUnarySum(lin int, col int, id string, op string, val interfaces.Expression) UnarySum {
	return UnarySum{lin, col, id, op, val}
}

func (p UnarySum) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}
