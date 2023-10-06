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
	var op1, op2, result environment.Value
	switch o.Operator {
	case "&&":
		{
			op1 = o.LeftExp.Execute(ast, env, gen)
			//add op1 labels
			for _, lvl := range op1.TrueLabel {
				gen.AddLabel(lvl.(string))
			}

			op2 = o.RightExp.Execute(ast, env, gen)

			result = environment.NewValue("", false, environment.BOOLEAN)
			result.TrueLabel = append(op2.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op1.FalseLabel, result.FalseLabel...)
			result.FalseLabel = append(op2.FalseLabel, result.FalseLabel...)
			return result
		}
	case "||":
		{
			op1 = o.LeftExp.Execute(ast, env, gen)

			for _, lvl := range op1.FalseLabel {
				gen.AddLabel(lvl.(string))
			}
			op2 = o.RightExp.Execute(ast, env, gen)

			result = environment.NewValue("", false, environment.BOOLEAN)

			result.TrueLabel = append(op1.TrueLabel, result.TrueLabel...)
			result.TrueLabel = append(op2.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op2.FalseLabel, result.FalseLabel...)
			return result
		}
	}
	gen.AddBr()
	return result
}
