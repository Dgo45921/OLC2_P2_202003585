package instructions

import (
	"PY1/environment"
	"PY1/generator"
)

type Break struct {
	Lin int
	Col int
}

func NewBreak(lin int, col int) Break {
	breakInstr := Break{lin, col}
	return breakInstr
}

func (p Break) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	return environment.Value{BreakFlag: true}
}
