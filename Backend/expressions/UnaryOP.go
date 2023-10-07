package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
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
	var op1, result environment.Value

	switch p.Operator {
	case "!":
		{
			op1 = p.Exp.Execute(ast, env, gen)
			if op1.Type == environment.BOOLEAN {
				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, op1.FalseLabel...)
				result.FalseLabel = append(result.FalseLabel, op1.TrueLabel...)
				return result
			} else {
				ast.SetError(p.Lin, p.Col, "ERROR: Type no compatible")
			}
		}
	case "-":
		{
			op1 = p.Exp.Execute(ast, env, gen)
			newTemp := gen.NewTemp()
			if op1.Type == environment.INTEGER {
				gen.AddExpression(newTemp, "0", op1.Value, "-")
				result = environment.NewValue(newTemp, true, environment.INTEGER)
				return result
			} else if op1.Type == environment.FLOAT {
				gen.AddExpression(newTemp, "0", op1.Value, "-")
				result = environment.NewValue(newTemp, true, environment.FLOAT)
				return result
			} else {
				fmt.Println("ERROR: tipo no compatible -")
			}
		}

	}
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
