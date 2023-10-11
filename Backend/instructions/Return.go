package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type Return struct {
	Lin int
	Col int
	Exp interface{}
}

func NewReturn(lin int, col int, exp interface{}) Return {
	breakInstr := Return{lin, col, exp}
	return breakInstr
}


func (p Return) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	if _, isBreak := p.Exp.(interfaces.Expression); isBreak {
		return p.Exp.(interfaces.Expression).Execute(ast, env, gen)

	}
	return environment.Value{}
}
