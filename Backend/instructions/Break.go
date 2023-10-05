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

func (p Break) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	if !env.(environment.Environment).InsideLoop() {
		ast.SetError(p.Lin, p.Col, "sentencia break debe de estar dentro de un ciclo")
	}
	return p
}
