package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type ManualMatrixDef struct {
	Lin   int
	Col   int
	Value interface{}
}

func NewManualMatrixDef(lin int, col int, val interface{}) ManualMatrixDef {
	exp := ManualMatrixDef{lin, col, val}
	return exp
}

func (p ManualMatrixDef) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value

	return result
}
