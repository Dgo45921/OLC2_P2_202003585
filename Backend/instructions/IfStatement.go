package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
	"strings"
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
	retorno := gen.NewLabel()                      //salida
	//*****************************************add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones if
	for _, inst := range p.TrueBlock {
		result = inst.(interfaces.Instruction).Execute(ast, newEnv, gen)

		//comprobar si es brak
		if result.BreakFlag {
			gen.AddGoto(gen.BreakLabel)
			result.BreakFlag = false
		} else if result.ContinueFlag {
			gen.AddGoto(gen.ContinueLabel)
			result.ContinueFlag = false
		} else if result.ReturnFlag {
			gen.AddGoto(gen.ReturnLabel)
			result.ReturnFlag = false
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

	//ELSE IF
	if len(p.ElseIfBlock) > 0 {
		//instrucciones elseif
		if gen.Flag {
			gen.Flag = false
			for _, s := range p.ElseIfBlock {
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
					} else if result.ReturnFlag {
						gen.AddGoto(gen.ReturnLabel)
						result.ReturnFlag = false
					}
					//out lbls
					for _, lvl := range result.OutLabel {
						OutLvls = append(OutLvls, lvl)
					}

				}
			}
		} else {
			gen.Auxlvl = retorno
			for _, s := range p.ElseIfBlock {
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
					} else if result.ReturnFlag {
						gen.AddGoto(gen.ReturnLabel)
						result.ReturnFlag = false
					}
					//out lbls
					for _, lvl := range result.OutLabel {
						OutLvls = append(OutLvls, lvl)
					}

				}
			}
		}
	}

	if len(p.ElseBlock) > 0 {
		//instrucciones elseif
		for _, s := range p.ElseBlock {
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
				} else if result.ReturnFlag {
					gen.AddGoto(gen.ReturnLabel)
					result.ReturnFlag = false
				}
				//out lbls
				for _, lvl := range result.OutLabel {
					OutLvls = append(OutLvls, lvl)
				}

			}
		}
		gen.Flag = false
	}

	gen.AddLabel(retorno)
	OutLvls = append(OutLvls, retorno)
	result.OutLabel = DeepCopyArray(OutLvls).([]interface{})
	return result
}

func (p If) Execute2(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("--- IF ---")
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
	for _, inst := range p.TrueBlock {
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

	// adding false labels
	for _, lvl := range condicion.FalseLabel {
		result.FalseLabel = append(result.FalseLabel, lvl.(string))
	}

	OutLvls = append(OutLvls, retorno)
	result.OutLabel = DeepCopyArray(OutLvls).([]interface{})
	return result
}
