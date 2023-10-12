package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
	"strings"
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
	gen.AddComment("Generando While")
	var condicion, result environment.Value
	//etiqueta de retorno
	RetLvl := gen.NewLabel()
	gen.AddLabel(RetLvl)
	//ejecutando expresion (if)
	condicion = p.Condition.Execute(ast, env, gen)
	//******************** add break & continue lvls
	gen.AddContinue(RetLvl)
	gen.AddBreak(condicion.FalseLabel[0].(string))
	//add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones while
	for _, s := range p.insBlock {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
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
			//agregando etiquetas de salida
			for _, lvl := range result.OutLabel {
				gen.AddLabel(lvl.(string))
			}
		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			result = s.(interfaces.Expression).Execute(ast, env, gen)
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
			//agregando etiquetas de salida
			for _, lvl := range result.OutLabel {
				gen.AddLabel(lvl.(string))
			}
		} else {
			fmt.Println("Error en bloque")
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
