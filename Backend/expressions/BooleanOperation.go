package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type BooleanOperation struct {
	Lin      int
	Col      int
	LeftExp  interfaces.Expression
	Operator string
	RightExp interfaces.Expression
}

func NewBooleanOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) BooleanOperation {
	exp := BooleanOperation{Lin: lin, Col: col, LeftExp: Op1, Operator: Operador, RightExp: Op2}
	return exp
}

func (o BooleanOperation) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
