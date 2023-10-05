package instructions

import (
	"PY1/environment"
	"PY1/generator"
)

type VecDec struct {
	Lin     int
	Col     int
	Id      string
	Type    string
	DefType interface{}
	Exp     interface{}
}

func NewVecDec(lin int, col int, id string, tyype string, deftype interface{}, exp interface{}) VecDec {
	NewVecDeclaration := VecDec{lin, col, id, tyype, deftype, exp}
	return NewVecDeclaration
}

func (p VecDec) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}
