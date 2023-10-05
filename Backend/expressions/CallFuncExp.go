package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type CallFuncExp struct {
	Lin        int
	Col        int
	Id         string
	Parameters []environment.FuncArg
}

func NewCallFuncExp(lin int, col int, val string, param []environment.FuncArg) CallFuncExp {
	exp := CallFuncExp{lin, col, val, param}
	return exp
}

func (p CallFuncExp) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
