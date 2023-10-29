package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type CountVector struct {
	Lin int
	Col int
	Id  string
}

func NewCountVector(lin int, col int, val string) CountVector {
	exp := CountVector{lin, col, val}
	return exp
}

func (p CountVector) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	var tempArray, result environment.Value

	prueba := VariableAccess{ID: p.Id}

	tempArray = prueba.Execute(ast, env, gen)

	//llamada
	newTmp := gen.NewTemp()

	tmp := gen.NewTemp()
	gen.AddGetHeap(tmp, "(int)"+tempArray.Value)

	gen.AddExpression(newTmp, tempArray.Value, "0", "+")
	gen.AddExpression(newTmp, newTmp, "0", "+")
	newTmp2 := gen.NewTemp()
	gen.AddGetHeap(newTmp2, "(int)"+newTmp)

	result = environment.Value{
		Value:        newTmp2,
		IsTemp:       true,
		Type:         getInsideType(tempArray.Type),
		TrueLabel:    nil,
		FalseLabel:   nil,
		OutLabel:     nil,
		IntValue:     0,
		FloatValue:   0,
		BreakFlag:    false,
		ContinueFlag: false,
	}
	return result

}
