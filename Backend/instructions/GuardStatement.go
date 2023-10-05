package instructions

import (
	"PY1/environment"
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

func (p Guard) Execute(ast *environment.AST, env interface{}) interface{} {
	var expResult = p.Condition.Execute(ast, env)
	if expResult.Type != environment.BOOLEAN {
		ast.SetError(p.Lin, p.Col, "La expresion del guard debe ser booleana")
		return nil
	}

	if expResult.Value.(bool) == false {
		var newEnv = environment.NewEnvironment(env, environment.GUARD)
		for _, inst := range p.FalseBlock {
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
		ast.SetError(p.Lin, p.Col, "Guard debe de terminar con continue, break o return")

	} else {
		return nil
	}

	return nil
}
