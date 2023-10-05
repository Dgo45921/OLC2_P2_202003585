package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type Case struct {
	Lin       int
	Col       int
	Condition interfaces.Expression
	insBlock  []interface{}
}

func NewCase(lin int, col int, condition interfaces.Expression, insBlock []interface{}) Case {
	return Case{lin, col, condition, insBlock}
}

func (p Case) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}