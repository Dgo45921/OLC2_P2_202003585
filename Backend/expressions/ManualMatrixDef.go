package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
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

func (p ManualMatrixDef) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	newArray := subtractOneFromElements(p.Value, ast, env)
	response := newArray.([]interface{})[0]

	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: response,
	}

}

func subtractOneFromElements(arr interface{}, ast *environment.AST, env interface{}) interface{} {
	switch arr.(type) {
	case []interface{}:
		result := make([]interface{}, len(arr.([]interface{})))
		for i, item := range arr.([]interface{}) {
			result[i] = subtractOneFromElements(item, ast, env)
		}
		return result
	case interfaces.Expression:
		return arr.(interfaces.Expression).Execute(ast, env).Value
	case []interfaces.Expression:
		result := make([]interface{}, len(arr.([]interface{})))
		for i, item := range arr.([]interface{}) {
			result[i] = item.(interfaces.Expression).Execute(ast, env).Value
		}
		return result
	default:
		return arr
	}
}
