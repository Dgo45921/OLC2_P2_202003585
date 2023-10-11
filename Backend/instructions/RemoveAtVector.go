package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type RemoveAtVector struct {
	Lin   int
	Col   int
	Id    string
	Index interfaces.Expression
}

func NewRemoveAtVector(lin int, col int, id string, index interfaces.Expression) RemoveAtVector {
	asig := RemoveAtVector{lin, col, id, index}
	return asig
}

func (p RemoveAtVector) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}