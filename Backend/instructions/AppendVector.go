package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type AppendVector struct {
	Lin        int
	Col        int
	Id         string
	Expression interfaces.Expression
}

func NewAppendVector(lin int, col int, id string, val interfaces.Expression) AppendVector {
	asig := AppendVector{lin, col, id, val}
	return asig
}

func (p AppendVector) Execute(ast *environment.AST, env interface{}) interface{} {

	if env.(environment.Environment).VariableExists(p.Id) {

		foundVar := env.(environment.Environment).FindVar(p.Id)
		if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR {
			if foundVar.Const {
				ast.SetError(p.Lin, p.Col, "No se puede modificar un vector constante")
				return nil

			}

			value := p.Expression.Execute(ast, env)

			if foundVar.Type == environment.VECTOR {
				if value.Type == environment.INTEGER {
					foundVar.Type = environment.VECTOR_INT
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.FLOAT {
					foundVar.Type = environment.VECTOR_FLOAT
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.BOOLEAN {
					foundVar.Type = environment.VECTOR_BOOLEAN
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.CHAR {
					foundVar.Type = environment.VECTOR_CHAR
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.STRING {
					foundVar.Type = environment.VECTOR_STRING
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.STRUCT_IMP {
					if foundVar.StructType == value.StructType {
						if _, isArray := foundVar.Value.([]interface{}); isArray {
							foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
						}
						env.(environment.Environment).UpdateVariable(p.Id, foundVar)
					} else {
						ast.SetError(p.Lin, p.Col, "append de struct distinto")

					}

				} else {
					ast.SetError(p.Lin, p.Col, "tipo de concatenacion incompatible")
				}
			} else {
				if value.Type == environment.INTEGER && foundVar.Type == environment.VECTOR_INT {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.FLOAT && foundVar.Type == environment.VECTOR_FLOAT {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.BOOLEAN && foundVar.Type == environment.VECTOR_BOOLEAN {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.STRING && foundVar.Type == environment.VECTOR_STRING {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.CHAR && foundVar.Type == environment.VECTOR_CHAR {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else if value.Type == environment.STRUCT_IMP && foundVar.Type == environment.VECTOR_STRUCT {
					if foundVar.StructType == value.StructType {
						if _, isArray := foundVar.Value.([]interface{}); isArray {
							foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
						}
						env.(environment.Environment).UpdateVariable(p.Id, foundVar)
					} else {
						ast.SetError(p.Lin, p.Col, "append de struct distinto")
					}

				} else {
					ast.SetError(p.Lin, p.Col, "tipo de concatenacion incompatible")
				}
			}

		} else {
			ast.SetError(p.Lin, p.Col, "funcion append solo funciona con vectores")
		}

		return nil
	} else if env.(environment.Environment).ReferenceExists(p.Id) {
		foundVar := env.(environment.Environment).FindReference(p.Id)
		if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR {
			if foundVar.Const {
				ast.SetError(p.Lin, p.Col, "No se puede modificar un vector constante")
				return nil

			}

			value := p.Expression.Execute(ast, env)

			if foundVar.Type == environment.VECTOR {
				if value.Type == environment.INTEGER {
					foundVar.Type = environment.VECTOR_INT
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.FLOAT {
					foundVar.Type = environment.VECTOR_FLOAT
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.BOOLEAN {
					foundVar.Type = environment.VECTOR_BOOLEAN
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.CHAR {
					foundVar.Type = environment.VECTOR_CHAR
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.STRING {
					foundVar.Type = environment.VECTOR_STRING
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.STRUCT_IMP {
					if foundVar.StructType == value.StructType {
						if _, isArray := foundVar.Value.([]interface{}); isArray {
							foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
						}
						env.(environment.Environment).UpdateVariable(p.Id, foundVar)
					} else {
						ast.SetError(p.Lin, p.Col, "append de struct distinto")
					}

				} else {
					ast.SetError(p.Lin, p.Col, "tipo de concatenacion incompatible")
				}
			} else {
				if value.Type == environment.INTEGER && foundVar.Type == environment.VECTOR_INT {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.FLOAT && foundVar.Type == environment.VECTOR_FLOAT {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.BOOLEAN && foundVar.Type == environment.VECTOR_BOOLEAN {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.STRING && foundVar.Type == environment.VECTOR_STRING {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.CHAR && foundVar.Type == environment.VECTOR_CHAR {
					if _, isArray := foundVar.Value.([]interface{}); isArray {
						foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
					}

				} else if value.Type == environment.STRUCT_IMP && foundVar.Type == environment.VECTOR_STRUCT {
					if foundVar.StructType == value.StructType {
						if _, isArray := foundVar.Value.([]interface{}); isArray {
							foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
						}
						env.(environment.Environment).UpdateVariable(p.Id, foundVar)
					} else {
						ast.SetError(p.Lin, p.Col, "append de struct distinto")
					}

				} else {
					ast.SetError(p.Lin, p.Col, "tipo de concatenacion incompatible")
				}
			}

		} else {
			ast.SetError(p.Lin, p.Col, "funcion append solo funciona con vectores")
		}

		return nil

	}

	return nil
}
