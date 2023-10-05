package instructions

import (
	"PY1/environment"
	"PY1/generator"
)

type RemoveLastVector struct {
	Lin int
	Col int
	Id  string
}

func NewRemoveLastVector(lin int, col int, id string) RemoveLastVector {
	asig := RemoveLastVector{lin, col, id}
	return asig
}

func (p RemoveLastVector) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}