package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type SelfModification struct {
	Lin int
	Col int
	Id  string
	NewValue interfaces.Expression
}

func NewSelfModification(lin int, col int, id string, newval interfaces.Expression) SelfModification {
	return SelfModification{lin, col, id, newval}
}

func (p SelfModification) Execute(ast *environment.AST, env interface{}) interface{} {
	return p
}
