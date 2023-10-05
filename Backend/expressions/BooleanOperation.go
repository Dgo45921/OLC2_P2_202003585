package expressions

import (
	"PY1/environment"
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

func (o BooleanOperation) Execute(ast *environment.AST, env interface{}) environment.Symbol {

	var op1, op2 environment.Symbol
	op1 = o.LeftExp.Execute(ast, env)
	op2 = o.RightExp.Execute(ast, env)

	switch o.Operator {

	case "||":
		{
			if (op1.Type == environment.BOOLEAN) && (op2.Type == environment.BOOLEAN) {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(bool) || op2.Value.(bool)}
			} else {
				ast.SetError(o.Lin, o.Col, "Operacion solo funciona con operandos booleanos")
			}
		}
	case "&&":
		{
			if (op1.Type == environment.BOOLEAN) && (op2.Type == environment.BOOLEAN) {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(bool) && op2.Value.(bool)}
			} else {
				ast.SetError(o.Lin, o.Col, "Operacion solo funciona con operandos booleanos")
			}
		}

	}

	var result interface{}
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.NULL, Value: result}
}
