package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"strconv"
)

type AppendVector struct {
	Lin        int
	Col        int
	Id         string
	Expression interfaces.Expression
}

func NewAppendVector(lin int, col int, id string, val interfaces.Expression) AppendVector {
	asig := AppendVector{lin, col, id, val}
	return asig
}

func (p AppendVector) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result, val environment.Value
	foundVector := env.(environment.Environment).FindVar(p.Id)

	size := len(foundVector.ArrayValues) + 1

	gen.AddComment("append array")
	newtmp1 := gen.NewTemp()
	newtmp2 := gen.NewTemp()
	gen.AddAssign(newtmp1, "H")
	gen.AddExpression(newtmp2, newtmp1, "1", "+")
	gen.AddSetHeap("(int)H", strconv.Itoa(size))
	gen.AddExpression("H", "H", strconv.Itoa(size+1), "+")

	foundVector.ArrayValues = append(foundVector.ArrayValues, p.Expression.(interfaces.Expression))

	for _, s := range foundVector.ArrayValues {
		val = s.(interfaces.Expression).Execute(ast, env, gen)
		gen.AddSetHeap("(int)"+newtmp2, val.Value)
		gen.AddExpression(newtmp2, newtmp2, "1", "+")
	}

	result = environment.Value{
		Value:        newtmp1,
		IsTemp:       true,
		Type:         getType(val.Type),
		TrueLabel:    nil,
		FalseLabel:   nil,
		OutLabel:     nil,
		IntValue:     0,
		FloatValue:   0,
		BreakFlag:    false,
		ContinueFlag: false,
	}

	newVar := env.(environment.Environment).AppendVector(p.Id, result.Type, size + 1, foundVector.ArrayValues)

	gen.AddSetStack(strconv.Itoa(newVar.Position), result.Value)
	gen.AddBr()

	return result

}
