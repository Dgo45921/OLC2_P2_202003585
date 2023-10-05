package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"fmt"
	"strconv"
)

type ArithmeticOperation struct {
	Lin      int
	Col      int
	LeftExp  interfaces.Expression
	Operator string
	RightExp interfaces.Expression
}

func NewArithmeticOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) ArithmeticOperation {
	exp := ArithmeticOperation{Lin: lin, Col: col, LeftExp: Op1, Operator: Operador, RightExp: Op2}
	return exp
}

func (o ArithmeticOperation) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	var dominante environment.TipoExpresion

	tablaDominante := [5][5]environment.TipoExpresion{
		//		INTEGER			FLOAT				STRING				BOOLEAN				NULL
		{environment.INTEGER, environment.FLOAT, environment.STRING, environment.BOOLEAN, environment.NULL},
		//FLOAT
		{environment.FLOAT, environment.FLOAT, environment.STRING, environment.NULL, environment.NULL},
		//STRING
		{environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.NULL},
		//BOOLEAN
		{environment.BOOLEAN, environment.NULL, environment.STRING, environment.BOOLEAN, environment.NULL},
		//NULL
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}

	var op1, op2 environment.Symbol
	op1 = o.LeftExp.Execute(ast, env)
	op2 = o.RightExp.Execute(ast, env)

	if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
		ast.SetPrint("Error: Tipo de operacion no valida!\n")
		return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.NULL, Value: nil}
	}

	switch o.Operator {
	case "+":
		{
			//validar tipo dominante
			dominante = tablaDominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: op1.Value.(int) + op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: val1 + val2}
			} else if dominante == environment.STRING {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: r1 + r2}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible realizar operacion")
			}
		}
	case "-":
		{
			dominante = tablaDominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: op1.Value.(int) - op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: val1 - val2}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible realizar operacion")
			}
		}
	case "*":
		{
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: op1.Value.(int) * op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: val1 * val2}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible realizar operacion")
			}
		}
	case "/":
		{
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				if op2.Value.(int) != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: op1.Value.(int) / op2.Value.(int)}
				} else {
					ast.SetError(o.Lin, o.Col, "No es posible dividir entre 0")
				}

			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				if val2 != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: val1 / val2}
				} else {
					ast.SetError(o.Lin, o.Col, "No es posible dividir entre 0")
				}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible dividir")
			}

		}

	case "%":
		{
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				if op2.Value.(int) != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: op1.Value.(int) % op2.Value.(int)}
				} else {
					ast.SetError(o.Lin, o.Col, "No es posible hacer el modulo con los operandos dados")
				}

			} else {
				ast.SetError(o.Lin, o.Col, "No es posible hacer el modulo con los operandos dados")
			}

		}

	}

	var result interface{}
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.NULL, Value: result}
}
