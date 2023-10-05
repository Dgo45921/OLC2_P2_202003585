package interfaces

import "PY1/environment"

type Expression interface {
	Execute(ast *environment.AST, env interface{}) environment.Symbol
}
