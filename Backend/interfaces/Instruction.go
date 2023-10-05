package interfaces

import "PY1/environment"

type Instruction interface {
	Execute(ast *environment.AST, env interface{}) interface{}
}
