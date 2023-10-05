package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type Asignation struct {
	Lin        int
	Col        int
	Id         string
	Expression interfaces.Expression
}

func NewAsignation(lin int, col int, id string, val interfaces.Expression) Asignation {
	asig := Asignation{lin, col, id, val}
	return asig
}

func (p Asignation) Execute(ast *environment.AST, env interface{}) interface{} {
	if env.(environment.Environment).VariableExists(p.Id) {
		foundVar := env.(environment.Environment).FindVar(p.Id)
		if !foundVar.Const {
			value := p.Expression.Execute(ast, env)

			if value.Type == environment.NULL {
				foundVar.Value = nil
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				return nil

			}

			if foundVar.Type == environment.NULL {
				foundVar.Type = value.Type
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				return nil
			}

			if value.Type == foundVar.Type {
				if value.Type == environment.STRUCT_IMP {
					if value.StructType == foundVar.StructType {
						foundVar.Value = value.Value
						env.(environment.Environment).UpdateVariable(p.Id, foundVar)
					} else {
						ast.SetError(p.Lin, p.Col, "Variable de struct no compatible con la asignacion dada")
						return nil
					}
				}

				if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.MATRIX_INT || foundVar.Type == environment.MATRIX_FLOAT || foundVar.Type == environment.MATRIX_STRING || foundVar.Type == environment.MATRIX_CHAR || foundVar.Type == environment.MATRIX_BOOLEAN || foundVar.Type == environment.VECTOR {
					foundVar.Value = DeepCopyArray(value.Value)
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
					return nil
				}
				foundVar.Value = value.Value
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)

			} else {
				ast.SetError(p.Lin, p.Col, "No se puede asignar valor: "+value.GetType()+" a variable "+foundVar.GetType())
				foundVar.Value = nil
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)

			}

		} else {
			ast.SetError(p.Lin, p.Col, "valor de una constante no puede ser cambiado")
		}

		return nil
	} else if env.(environment.Environment).ReferenceExists(p.Id) {
		foundVar := env.(environment.Environment).FindReference(p.Id)
		if !foundVar.Const {
			value := p.Expression.Execute(ast, env)

			if value.Type == environment.NULL {
				foundVar.Value = nil
				env.(environment.Environment).UpdateReference(p.Id, foundVar)
				return nil

			}

			if foundVar.Type == environment.NULL {
				foundVar.Type = value.Type
				env.(environment.Environment).UpdateReference(p.Id, foundVar)
				return nil
			}

			if value.Type == foundVar.Type {
				if value.Type == environment.STRUCT_IMP {
					if value.StructType == foundVar.StructType {
						foundVar.Value = value.Value
						env.(environment.Environment).UpdateReference(p.Id, foundVar)
					} else {
						ast.SetError(p.Lin, p.Col, "Variable de struct no compatible con la asignacion dada")
						return nil
					}
				}

				if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.MATRIX_INT || foundVar.Type == environment.MATRIX_FLOAT || foundVar.Type == environment.MATRIX_STRING || foundVar.Type == environment.MATRIX_CHAR || foundVar.Type == environment.MATRIX_BOOLEAN || foundVar.Type == environment.VECTOR {
					foundVar.Value = DeepCopyArray(value.Value)
					env.(environment.Environment).UpdateReference(p.Id, foundVar)
					return nil
				}
				foundVar.Value = value.Value
				env.(environment.Environment).UpdateReference(p.Id, foundVar)

			} else {
				ast.SetError(p.Lin, p.Col, "No se puede asignar valor: "+value.GetType()+" a variable "+foundVar.GetType())

				foundVar.Value = nil
				env.(environment.Environment).UpdateReference(p.Id, foundVar)

			}

		} else {
			ast.SetError(p.Lin, p.Col, "valor de una constante no puede ser cambiado")
		}

		return nil
	}

	return nil
}

func DeepCopyArray(source interface{}) interface{} {
	switch source := source.(type) {
	case []interface{}:
		copyArray := make([]interface{}, len(source))
		for i, val := range source {
			copyArray[i] = DeepCopyArray(val)
		}
		return copyArray
	default:
		return source
	}
}
