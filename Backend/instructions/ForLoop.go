package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type For struct {
	Lin      int
	Col      int
	Id       string
	Range    interfaces.Expression
	insBlock []interface{}
}

func NewFor(lin int, col int, id string, rangge interfaces.Expression, insBlock []interface{}) For {
	return For{lin, col, id, rangge, insBlock}
}

func (p For) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}