package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type ArithmeticOperation struct {
	Lin      int
	Col      int
	OpIzq    interfaces.Expression
	Operator string
	OpDer    interfaces.Expression
}

func NewArithmeticOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) ArithmeticOperation {
	exp := ArithmeticOperation{Lin: lin, Col: col, OpIzq: Op1, Operator: Operador, OpDer: Op2}
	return exp
}

func (o ArithmeticOperation) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
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

	var op1, op2, result environment.Value

	newTemp := gen.NewTemp()
	switch o.Operator {
	case "+":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			op2 = o.OpDer.Execute(ast, env, gen)
			//validar tipo dominante
			dominante = tablaDominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "+")
				result = environment.NewValue(newTemp, true, dominante)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible sumar")
			}
		}
	case "-":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			op2 = o.OpDer.Execute(ast, env, gen)
			dominante = tablaDominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "-")
				result = environment.NewValue(newTemp, true, dominante)
				result.IntValue = op1.IntValue - op2.IntValue
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible restar")
			}
		}
	case "*":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			op2 = o.OpDer.Execute(ast, env, gen)
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "*")
				result = environment.NewValue(newTemp, true, dominante)
				result.IntValue = op1.IntValue * op2.IntValue
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible multiplicar")
			}
		}
	case "/":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			op2 = o.OpDer.Execute(ast, env, gen)
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				lvl1 := gen.NewLabel()
				lvl2 := gen.NewLabel()

				gen.AddIf(op2.Value, "0", "!=", lvl1)
				gen.AddPrintf("c", "77")
				gen.AddPrintf("c", "97")
				gen.AddPrintf("c", "116")
				gen.AddPrintf("c", "104")
				gen.AddPrintf("c", "69")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "111")
				gen.AddPrintf("c", "114")
				gen.AddExpression(newTemp, "0", "", "")
				gen.AddGoto(lvl2)
				gen.AddLabel(lvl1)
				gen.AddExpression(newTemp, op1.Value, op2.Value, "/")
				gen.AddLabel(lvl2)
				result = environment.NewValue(newTemp, true, dominante)
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible dividir")
			}

		}


	}
	gen.AddBr()
	return environment.Value{}
}
