package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type AppendVector struct {
	Lin        int
	Col        int
	Id         string
	Expression interfaces.Expression
}

func NewAppendVector(lin int, col int, id string, val interfaces.Expression) AppendVector {
	asig := AppendVector{lin, col, id, val}
	return asig
}

func (p AppendVector) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
