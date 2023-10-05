package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type UnaryOp struct {
	Lin      int
	Col      int
	Exp      interfaces.Expression
	Operator string
}

func NewUnaryOperation(lin int, col int, Op1 interfaces.Expression, Operador string) UnaryOp {
	exp := UnaryOp{Lin: lin, Col: col, Exp: Op1, Operator: Operador}
	return exp
}

func (p UnaryOp) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
func returnString(val int) string {
	if val == 0 {
		return "INTEGER"

	} else if val == 1 {
		return "FLOAT"

	} else if val == 2 {
		return "STRING"

	} else if val == 3 {
		return "BOOLEAN"
	}

	return "NULL"
}
