package expressions

import (
	"PY1/environment"
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

func (o UnaryOp) Execute(ast *environment.AST, env interface{}) environment.Symbol {

	var op1 environment.Symbol
	op1 = o.Exp.Execute(ast, env)

	switch o.Operator {
	case "!":
		{
			if op1.Type == environment.BOOLEAN {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: !op1.Value.(bool)}

			} else {
				ast.SetError(o.Lin, o.Col, "No es restar con operador -=")
			}
		}
	case "-":
		{
			if op1.Type == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.INTEGER, Value: -1 * op1.Value.(int)}

			} else if op1.Type == environment.FLOAT {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.INTEGER, Value: -1 * op1.Value.(float64)}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible usar negativo en: "+returnString(int(op1.Type)))
			}

		}

	}

	var result interface{}
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.NULL, Value: result}
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
