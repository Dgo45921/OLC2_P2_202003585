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

	var result environment.Value

	VariableAccess{ID: p.Id}.Execute(ast, env, gen)

	//llamada
	newTmp := gen.NewTemp()
	lvl1 := gen.NewLabel()
	lvl2 := gen.NewLabel()
	lvl3 := gen.NewLabel()

	gen.AddGoto(lvl2)
	gen.AddLabel(lvl1)
	gen.AddPrintf("c", "66")
	gen.AddPrintf("c", "111")
	gen.AddPrintf("c", "117")
	gen.AddPrintf("c", "110")
	gen.AddPrintf("c", "100")
	gen.AddPrintf("c", "115")
	gen.AddPrintf("c", "69")
	gen.AddPrintf("c", "114")
	gen.AddPrintf("c", "114")
	gen.AddPrintf("c", "111")
	gen.AddPrintf("c", "114")
	gen.AddGoto(lvl3)
	gen.AddLabel(lvl2)

	newTmp2 := gen.NewTemp()
	gen.AddGetHeap(newTmp2, "(int)"+newTmp)
	gen.AddLabel(lvl3)

	result = environment.Value{
		Value:        newTmp2,
		IsTemp:       true,
		Type:         environment.INTEGER,
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
