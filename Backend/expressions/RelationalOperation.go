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

func (o RelationalOperation) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var op1, op2, result environment.Value
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

	switch o.Operator {
	case "<":
		{
			op1 = o.LeftExp.Execute(ast, env, gen)
			op2 = o.RightExp.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "<", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result

			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible comparar <")
			}
		}
	case ">":
		{
			op1 = o.LeftExp.Execute(ast, env, gen)
			op2 = o.RightExp.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, ">", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible comparar >")
			}
		}
	case "<=":
		{
			op1 = o.LeftExp.Execute(ast, env, gen)
			op2 = o.RightExp.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "<=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible comparar <=")
			}
		}
	case ">=":
		{
			op1 = o.LeftExp.Execute(ast, env, gen)
			op2 = o.RightExp.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, ">=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible comparar >=")
			}
		}
	case "==":
		{
			op1 = o.LeftExp.Execute(ast, env, gen)
			op2 = o.RightExp.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "==", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible comparar ==")
			}
		}
	case "!=":
		{
			op1 = o.LeftExp.Execute(ast, env, gen)
			op2 = o.RightExp.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "!=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible comparar !=")
			}
		}
	}
	return result
}
