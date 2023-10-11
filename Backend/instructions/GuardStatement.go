package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type Guard struct {
	Lin        int
	Col        int
	Condition  interfaces.Expression
	FalseBlock []interface{}
}

func NewGuard(lin int, col int, condition interfaces.Expression, falseb []interface{}) Guard {
	return Guard{lin, col, condition, falseb}
}

func (p Guard) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
