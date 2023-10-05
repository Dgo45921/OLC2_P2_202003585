package instructions

import (
	"PY1/environment"
	"PY1/generator"
)

type CallFuncInst struct {
	Lin        int
	Col        int
	Id         string
	Parameters []environment.FuncArg
}

func NewCallFuncInst(lin int, col int, val string, param []environment.FuncArg) CallFuncInst {
	exp := CallFuncInst{lin, col, val, param}
	return exp
}

func (p CallFuncInst) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}
