package instructions

import (
	"PY1/environment"
	"PY1/generator"
)

type Continue struct {
	Lin int
	Col int
}

func NewContinue(lin int, col int) Continue {
	breakInstr := Continue{lin, col}
	return breakInstr
}

func (p Continue) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value{
	if !env.(environment.Environment).InsideLoop() {
		ast.SetError(p.Lin, p.Col, "sentencia continue debe de estar dentro de un ciclo")
	}
	return environment.Value{}
}
