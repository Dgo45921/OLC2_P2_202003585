package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"strconv"
)

type VariableAccess struct {
	ID string
}

func NewVariableAccess(id string) VariableAccess {
	exp := VariableAccess{id}
	return exp
}

func (p VariableAccess) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("Llamando una variable")
	var result environment.Value
	retSym := env.(environment.Environment).FindVar(p.ID)
	newTemp := gen.NewTemp()
	newTemp2 := gen.NewTemp()
	if gen.MainCode {
		gen.AddGetStack(newTemp2, strconv.Itoa(retSym.Position))
	} else {
		gen.AddExpression(newTemp, "P", strconv.Itoa(retSym.Position), "+")
		gen.AddGetStack(newTemp2, "(int)"+newTemp)
	}

	if retSym.Type == environment.BOOLEAN {
		trueLabel := gen.NewLabel()
		falseLabel := gen.NewLabel()
		gen.AddIf(newTemp2, "1", "==", trueLabel)
		gen.AddGoto(falseLabel)
		result = environment.NewValue("", false, environment.BOOLEAN)
		result.TrueLabel = append(result.TrueLabel, trueLabel)
		result.FalseLabel = append(result.FalseLabel, falseLabel)
	} else {
		result = environment.NewValue(newTemp2, true, retSym.Type)
		result.Dimentions = retSym.Dimentions
	}
	return result
}
