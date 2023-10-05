package expressions

import (
	"PY1/environment"
)

type SelfAccess struct {
	Lin int
	Col int
	ID  string
}

func NewSelfAccess(lin int, col int, id string) SelfAccess {
	structaccess := SelfAccess{lin, col, id}
	return structaccess
}

func (p SelfAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {

	return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: p.ID}
}
