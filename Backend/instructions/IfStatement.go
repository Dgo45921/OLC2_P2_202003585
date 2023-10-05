package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type If struct {
	Lin         int
	Col         int
	Condition   interfaces.Expression
	TrueBlock   []interface{}
	ElseIfBlock []interface{}
	ElseBlock   []interface{}
}

func NewIf(lin int, col int, condition interfaces.Expression, trueb []interface{}, elif []interface{}, elsse []interface{}) If {
	return If{lin, col, condition, trueb, elif, elsse}
}

func (p If) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}
