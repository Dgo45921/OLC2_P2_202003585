package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type VariableAccess struct {
	ID string
}

func NewVariableAccess(id string) VariableAccess {
	exp := VariableAccess{id}
	return exp
}

func (p VariableAccess) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
