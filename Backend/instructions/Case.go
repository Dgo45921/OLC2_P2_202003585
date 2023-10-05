package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type Case struct {
	Lin       int
	Col       int
	Condition interfaces.Expression
	insBlock  []interface{}
}

func NewCase(lin int, col int, condition interfaces.Expression, insBlock []interface{}) Case {
	return Case{lin, col, condition, insBlock}
}

func (p Case) Execute(ast *environment.AST, env interface{}) interface{} {
	var newEnv = environment.NewEnvironment(env, environment.CASE)
	for _, inst := range p.insBlock {
		var response = inst.(interfaces.Instruction).Execute(ast, newEnv)
		if response != nil {
			if rep, isContinue := response.(Continue); isContinue {
				return rep
			}
			if _, isContinue := response.(Break); isContinue {
				break
			}
		}

	}

	return nil
}
