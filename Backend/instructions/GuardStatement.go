package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type Guard struct {
	Lin        int
	Col        int
	Condition  interfaces.Expression
	FalseBlock []interface{}
}

func NewGuard(lin int, col int, condition interfaces.Expression, falseb []interface{}) Guard {
	return Guard{lin, col, condition, falseb}
}

func (p Guard) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("--- GUARD ---")
	var condicion, result environment.Value
	var newEnv = environment.NewEnvironment(env, environment.IF)
	var OutLvls []interface{}
	condicion = p.Condition.Execute(ast, env, gen) //imprime el if de operacion
	retorno := gen.NewLabel()                      //salida
	//*****************************************add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones if
	for _, inst := range p.FalseBlock {
		result = inst.(interfaces.Instruction).Execute(ast, newEnv, gen)

		//comprobar si es brak
		if result.BreakFlag {
			gen.AddGoto(gen.BreakLabel)
			result.BreakFlag = false
		} else if result.ContinueFlag {
			gen.AddGoto(gen.ContinueLabel)
			result.ContinueFlag = false
		}

		//out lbls

		for _, lvl := range result.OutLabel {
			OutLvls = append(OutLvls, lvl)
		}

	}

	if !gen.Flag {
		gen.AddGoto(retorno)
	} else {
		gen.AddGoto(gen.Auxlvl)
	}

	// adding false labels
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}

	gen.AddLabel(retorno)
	OutLvls = append(OutLvls, retorno)
	result.OutLabel = DeepCopyArray(OutLvls).([]interface{})
	return result
}
