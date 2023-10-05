package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type RelationalOperation struct {
	Lin      int
	Col      int
	LeftExp  interfaces.Expression
	Operator string
	RightExp interfaces.Expression
}

func NewRelationalOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) RelationalOperation {
	exp := RelationalOperation{Lin: lin, Col: col, LeftExp: Op1, Operator: Operador, RightExp: Op2}
	return exp
}

func (p RelationalOperation) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
