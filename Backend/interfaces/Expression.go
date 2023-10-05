package interfaces

import (
	"PY1/environment"
	"PY1/generator"
)

type Expression interface {
	Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value
}