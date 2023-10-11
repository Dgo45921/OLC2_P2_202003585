package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type While struct {
	Lin       int
	Col       int
	Condition interfaces.Expression
	insBlock  []interface{}
}

func NewWhile(lin int, col int, condition interfaces.Expression, insBlock []interface{}) While {
	return While{lin, col, condition, insBlock}
}

func (p While) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
