package instructions

import (
	"PY1/environment"
	"PY1/generator"
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

func (p SelfModification) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value{
	return environment.Value{}
}
