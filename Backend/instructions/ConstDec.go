package instructions

import (
	"PY1/environment"
	"PY1/generator"
)

type ConstDec struct {
	Lin        int
	Col        int
	Id         string
	Type       interface{}
	Expression interface{}
}

func NewConstDec(lin int, col int, id string, tyype interface{}, val interface{}) ConstDec {
	NewConstDeclaration := ConstDec{lin, col, id, tyype, val}
	return NewConstDeclaration
}

func (p ConstDec) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}