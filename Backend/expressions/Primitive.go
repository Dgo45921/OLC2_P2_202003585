package expressions

import (
	"PY1/environment"
)

type Primitive struct {
	Lin   int
	Col   int
	Value interface{}
	Type  environment.TipoExpresion
}

func NewPrimitive(lin int, col int, valor interface{}, tipo environment.TipoExpresion) Primitive {
	exp := Primitive{lin, col, valor, tipo}
	return exp
}

func (p Primitive) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Type:  p.Type,
		Value: p.Value,
	}
}
