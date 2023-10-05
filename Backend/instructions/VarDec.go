package instructions

import (
	"PY1/environment"
	"PY1/generator"
)

type VarDec struct {
	Lin        int
	Col        int
	Id         string
	Type       interface{}
	Expression interface{}
}

func NewVarDec(lin int, col int, id string, tyype interface{}, val interface{}) VarDec {
	NewVarDeclaration := VarDec{lin, col, id, tyype, val}
	return NewVarDeclaration
}

func (p VarDec) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}
