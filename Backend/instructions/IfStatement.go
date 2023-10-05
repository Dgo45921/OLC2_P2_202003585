package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
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

func (p If) Execute(ast *environment.AST, env interface{}) interface{} {
	var shouldExecuteElse = false
	var expResult = p.Condition.Execute(ast, env)
	if expResult.Type != environment.BOOLEAN {
		ast.SetError(p.Lin, p.Col, "La expresion dentro del if debe ser booleana")
		return nil
	}

	if expResult.Value.(bool) == true {
		var newEnv = environment.NewEnvironment(env, environment.IF)
		for _, inst := range p.TrueBlock {
			var response = inst.(interfaces.Instruction).Execute(ast, newEnv)
			if response != nil {
				if _, isBreak := response.(Break); isBreak {
					return response
				} else if _, isContinue := response.(Continue); isContinue {
					return response
				} else if _, isReturn := response.(environment.Symbol); isReturn {
					return response

				}
			}
		}
	} else {
		shouldExecuteElse = true

		if len(p.ElseIfBlock) > 0 {
			var newEnv = environment.NewEnvironment(env, environment.ELSEIF)
			for _, inst := range p.ElseIfBlock {
				if instruction, isInstruction := inst.(If); isInstruction {
					var condResult = inst.(If).Condition.Execute(ast, newEnv)
					if condResult.Value.(bool) {
						shouldExecuteElse = false
						var response = instruction.Execute(ast, newEnv)
						if response != nil {
							if _, isBreak := response.(Break); isBreak {
								return response
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

		if shouldExecuteElse {

			if len(p.ElseBlock) > 0 {
				var newEnv = environment.NewEnvironment(env, environment.ElSE)
				for _, inst := range p.ElseBlock {
					if instruction, isInstruction := inst.(interfaces.Instruction); isInstruction {
						var response = instruction.(interfaces.Instruction).Execute(ast, newEnv)
						if response != nil {
							if _, isBreak := response.(Break); isBreak {
								return response
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

	}

	return nil
}
