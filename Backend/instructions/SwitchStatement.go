package instructions

import (
	"PY1/environment"
	"PY1/expressions"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
	"strings"
)

type Switch struct {
	Lin          int
	Col          int
	Condition    interfaces.Expression
	CasesList    []interface{}
	DefaultBlock []interface{}
}

func NewSwitch(lin int, col int, condition interfaces.Expression, casesList []interface{}, defaultstmt []interface{}) Switch {
	return Switch{lin, col, condition, casesList, defaultstmt}
}

func (p Switch) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("--- SWITCH ---")
	var result environment.Value
	var OutLvls []interface{}
	retorno := gen.NewLabel() //salida
	//*****************************************add true labels

	//ELSE IF
	if len(p.CasesList) > 0 {
		//instrucciones elseif
		if gen.Flag {
			gen.Flag = false
			for _, s := range p.CasesList {
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
					//out lbls
					for _, lvl := range result.OutLabel {
						OutLvls = append(OutLvls, lvl)
					}

				}
			}
		} else {
			gen.Auxlvl = retorno
			for _, s := range p.CasesList {
				if strings.Contains(fmt.Sprintf("%T", s), "instructions") {

					targetValue := expressions.RelationalOperation{
						Lin:      0,
						Col:      0,
						LeftExp:  p.Condition,
						Operator: "==",
						RightExp: s.(Case).Condition,
					}

					resultt := targetValue.Execute(ast, env, gen)
					for _, lvl := range resultt.TrueLabel {
						gen.AddLabel(lvl.(string))
					}
					for _, ins := range s.(Case).insBlock {
						res2 := ins.(interfaces.Instruction).Execute(ast, env, gen)
						if res2.BreakFlag {
							gen.AddGoto(gen.BreakLabel)
							res2.BreakFlag = false
						}
						//comprobar si es continue
						if res2.ContinueFlag {
							gen.AddGoto(gen.ContinueLabel)
							res2.ContinueFlag = false
						}
						//out lbls
						for _, lvl := range res2.OutLabel {
							OutLvls = append(OutLvls, lvl)
						}

					}

					if !gen.Flag {
						gen.AddGoto(retorno)
					} else {
						gen.AddGoto(gen.Auxlvl)
					}

					for _, lvl := range resultt.FalseLabel {
						gen.AddLabel(lvl.(string))
					}

				}
			}

			//for _, s := range p.CasesList {
			//	if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			//		result = s.(Case).Execute(ast, env, gen)
			//		targetValue := s.(Case).Condition.Execute(ast, env, gen)
			//		caseLabel := gen.NewLabel()
			//		gen.AddIf(condicion.Value, targetValue.Value, "==", caseLabel)
			//
			//		gen.AddLabel(caseLabel)
			//		for _, ci := range s.(Case).insBlock {
			//			res2 := ci.(interfaces.Instruction).Execute(ast, env, gen)
			//
			//			//comprobar si es brak
			//			if res2.BreakFlag {
			//				gen.AddGoto(gen.BreakLabel)
			//				res2.BreakFlag = false
			//			}
			//			//comprobar si es continue
			//			if res2.ContinueFlag {
			//				gen.AddGoto(gen.ContinueLabel)
			//				res2.ContinueFlag = false
			//			}
			//			//out lbls
			//			for _, lvl := range res2.OutLabel {
			//				OutLvls = append(OutLvls, lvl)
			//			}
			//
			//		}
			//		gen.AddGoto(retorno)
			//
			//	}
			//}
		}
	}

	if len(p.DefaultBlock) > 0 {
		//instrucciones elseif
		for _, s := range p.DefaultBlock {
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
