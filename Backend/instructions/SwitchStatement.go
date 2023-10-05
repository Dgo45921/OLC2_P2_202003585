package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
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

func (p Switch) Execute(ast *environment.AST, env interface{}) interface{} {
	var shouldExecuteDefault = false
	var targetValue = p.Condition.Execute(ast, env).Value
	shouldExecuteDefault = true

	if len(p.CasesList) > 0 {
		for _, inst := range p.CasesList {
			if instruction, isCase := inst.(Case); isCase {
				var newEnv = environment.NewEnvironment(env, environment.CASE)
				var condResult = inst.(Case).Condition.Execute(ast, newEnv)
				if condResult.Value == targetValue {
					shouldExecuteDefault = false
					var response = instruction.Execute(ast, newEnv)
					if response != nil {
						if _, isBreak := response.(Break); isBreak {
							break
						} else if _, isContinue := response.(Continue); isContinue {
							return response
						} else if _, isReturn := response.(environment.Symbol); isReturn {
							return response

						}
					}
					break
				}

			}
		}

	}

	if shouldExecuteDefault {

		if len(p.DefaultBlock) > 0 {
			var newEnv = environment.NewEnvironment(env, environment.DEFAULT)
			for _, inst := range p.DefaultBlock {
				if instruction, isInstruction := inst.(interfaces.Instruction); isInstruction {
					var response = instruction.(interfaces.Instruction).Execute(ast, newEnv)
					if response != nil {
						if _, isBreak := response.(Break); isBreak {
							break
						} else if _, isContinue := response.(Continue); isContinue {
							return response
						} else if _, isReturn := response.(environment.Symbol); isReturn {
							return response

						}
					}

				}
			}

		}

	}

	return nil
}
