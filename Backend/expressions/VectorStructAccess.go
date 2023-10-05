package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type VectorStructAccess struct {
	Lin      int
	Col      int
	IDvector string
	Index    []interface{}
	Accesses []string
}

func NewVectorStructAccess(lin int, col int, id string, index []interface{}, acc []string) VectorStructAccess {
	exp := VectorStructAccess{lin, col, id, index, acc}
	return exp
}

func (p VectorStructAccess) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
