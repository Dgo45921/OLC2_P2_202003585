package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
)

type ManualMatrixDef struct {
	Lin   int
	Col   int
	Value interface{}
}

func NewManualMatrixDef(lin int, col int, val interface{}) ManualMatrixDef {
	exp := ManualMatrixDef{lin, col, val}
	return exp
}

func (p ManualMatrixDef) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	newArray := subtractOneFromElements(p.Value, ast, env, gen)
	response := newArray.([]interface{})[0]
	fmt.Println(response)
	var result environment.Value
	return result
}

func subtractOneFromElements(arr interface{}, ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	switch arr.(type) {
	case []interface{}:
		result := make([]interface{}, len(arr.([]interface{})))
		for i, item := range arr.([]interface{}) {
			result[i] = subtractOneFromElements(item, ast, env, gen)
		}
		return result
	case interfaces.Expression:
		return arr.(interfaces.Expression)
	case []interfaces.Expression:
		result := make([]interface{}, len(arr.([]interface{})))
		for i, item := range arr.([]interface{}) {
			result[i] = item.(interfaces.Expression)
		}
		return result
	default:
		return arr
	}
}
