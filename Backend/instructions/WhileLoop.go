package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type While struct {
	Lin       int
	Col       int
	Condition interfaces.Expression
	insBlock  []interface{}
}

func NewWhile(lin int, col int, condition interfaces.Expression, insBlock []interface{}) While {
	return While{lin, col, condition, insBlock}
}

func (p While) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("---WHILE---")
	var condicion, result environment.Value
	//etiqueta de retorno
	RetLvl := gen.NewLabel()
	gen.AddLabel(RetLvl)
	//ejecutando expresion (if)
	condicion = p.Condition.Execute(ast, env, gen)
	gen.AddContinue(RetLvl)
	gen.AddBreak(condicion.FalseLabel[0].(string))
	//******************** add break & continue lvls
	//add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones while
	for _, s := range p.insBlock {
		result = s.(interfaces.Instruction).Execute(ast, env, gen)
		//comprobar si es brak
		if result.BreakFlag {
			gen.AddGoto(gen.BreakLabel)
			result.BreakFlag = false
		}
		//comprobar si es continue
		if result.ContinueFlag {
			gen.AddGoto(gen.ContinueLabel)
			result.ContinueFlag = false
		}

	}

	//retorno
	gen.AddGoto(RetLvl)
	//add false labels
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}
	return result
}
