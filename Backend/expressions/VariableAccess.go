package expressions

import (
	"PY1/environment"
)

type VariableAccess struct {
	ID string
}

func NewVariableAccess(id string) VariableAccess {
	exp := VariableAccess{id}
	return exp
}

func (p VariableAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	if env.(environment.Environment).VariableExists(p.ID) {
		result := env.(environment.Environment).FindVar(p.ID)
		return result
	} else if env.(environment.Environment).ReferenceExists(p.ID) {
		result := env.(environment.Environment).FindReference(p.ID)
		return result
	}
	return environment.Symbol{
		Lin:   0,
		Col:   0,
		Value: nil,
	}

}
