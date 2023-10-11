package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
)

type If struct {
	Lin         int
	Col         int
	Condition   interfaces.Expression
	TrueBlock   []interface{}
	ElseIfBlock []interface{}
	ElseBlock   []interface{}
}

func NewIf(lin int, col int, condition interfaces.Expression, trueb []interface{}, elif []interface{}, elsse []interface{}) If {
	return If{lin, col, condition, trueb, elif, elsse}
}

func (p If) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("--- IF ---")
	var condicion, result environment.Value
	var newEnv = environment.NewEnvironment(env, environment.IF)
	var OutLvls []interface{}
	condicion = p.Condition.Execute(ast, env, gen) //imprime el if de operacion
	newLabel := gen.NewLabel()                     //salida
	//*****************************************add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones if
	for _, inst := range p.TrueBlock {
		result = inst.(interfaces.Instruction).Execute(ast, newEnv, gen)
		fmt.Println(result)
		//comprobar si es brak
		//if result.BreakFlag {
		//	gen.AddGoto(gen.BreakLabel)
		//	result.BreakFlag = false
		//}
		////comprobar si es continue
		//if result.ContinueFlag {
		//	gen.AddGoto(gen.ContinueLabel)
		//	result.ContinueFlag = false
		//}
		//out lbls

		for _, lvl := range result.OutLabel {
			OutLvls = append(OutLvls, lvl)
		}



	}

	gen.AddGoto(newLabel)
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}

	OutLvls = append(OutLvls, newLabel)
	result.OutLabel = OutLvls
	return result
}
