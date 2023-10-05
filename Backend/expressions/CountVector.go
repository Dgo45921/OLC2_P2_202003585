package expressions

import (
	"PY1/environment"
)

type CountVector struct {
	Lin int
	Col int
	Id  string
}

func NewCountVector(lin int, col int, val string) CountVector {
	exp := CountVector{lin, col, val}
	return exp
}

func (p CountVector) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	if env.(environment.Environment).VariableExists(p.Id) {
		foundVar := env.(environment.Environment).FindVar(p.Id)
		if _, isArray := foundVar.Value.([]interface{}); isArray {
			if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT {
				long := len(foundVar.Value.([]interface{}))
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: long,
					Type:  environment.INTEGER,
					Const: false,
				}

			}

		}
		ast.SetError(p.Lin, p.Col, "la funcion count solo funciona con vectores")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	} else if env.(environment.Environment).ReferenceExists(p.Id) {
		foundVar := env.(environment.Environment).FindReference(p.Id)
		if _, isArray := foundVar.Value.([]interface{}); isArray {
			if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT {
				long := len(foundVar.Value.([]interface{}))
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: long,
					Type:  environment.INTEGER,
					Const: false,
				}

			}

		}
		ast.SetError(p.Lin, p.Col, "la funcion count solo funciona con vectores")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}

	}
	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: nil,
	}
}
