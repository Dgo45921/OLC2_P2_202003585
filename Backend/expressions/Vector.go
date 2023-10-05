package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type Vector struct {
	Lin   int
	Col   int
	Value []interface{}
}

func NewVector(lin int, col int, val []interface{}) Vector {
	exp := Vector{lin, col, val}
	return exp
}

func (p Vector) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	if len(p.Value) == 0 {
		var emptyArray []interface{} = []interface{}{}
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: emptyArray,
			Type:  environment.VECTOR,
			Const: false,
		}
	}

	if validVector(ast, env, p) {
		var firstType = environment.INTEGER

		var valuesArray []interface{} = []interface{}{}
		for _, inst := range p.Value {
			var response = inst.(interfaces.Expression).Execute(ast, env)
			firstType = response.Type
			valuesArray = append(valuesArray, response.Value)
		}

		var vectype = environment.VECTOR_INT
		if firstType == environment.INTEGER {
			vectype = environment.VECTOR_INT
		} else if firstType == environment.FLOAT {
			vectype = environment.VECTOR_FLOAT
		} else if firstType == environment.BOOLEAN {
			vectype = environment.VECTOR_BOOLEAN
		} else if firstType == environment.STRING {
			vectype = environment.VECTOR_STRING
		} else if firstType == environment.CHAR {
			vectype = environment.VECTOR_CHAR
		} else if firstType == environment.STRUCT_IMP {
			vectype = environment.VECTOR_STRUCT
		}

		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: valuesArray,
			Type:  vectype,
			Const: false,
		}

	} else {
		ast.SetError(p.Lin, p.Col, "Vector no valido, debe de ser solo de un tipo")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	}

}

func validVector(ast *environment.AST, env interface{}, vector Vector) bool {
	var firstType = environment.INTEGER

	if len(vector.Value) == 0 {
		return true
	}

	if len(vector.Value) > 0 {

		if _, isExp := vector.Value[0].(interfaces.Expression); isExp {
			firstType = vector.Value[0].(interfaces.Expression).Execute(ast, env).Type
		}

	}

	for _, inst := range vector.Value {
		if _, isExp := inst.(interfaces.Expression); isExp {
			var response = inst.(interfaces.Expression).Execute(ast, env)
			if response.Type != firstType {
				return false
			}
		}

	}

	if firstType == environment.STRUCT_IMP {
		var structType = vector.Value[0].(interfaces.Expression).Execute(ast, env).StructType
		for _, inst := range vector.Value {
			if _, isExp := inst.(interfaces.Expression); isExp {
				var response = inst.(interfaces.Expression).Execute(ast, env)
				if response.StructType != structType {
					return false
				}

			}

		}
	}

	return true
}
