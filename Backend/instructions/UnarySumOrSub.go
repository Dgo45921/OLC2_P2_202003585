package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type UnarySum struct {
	Lin        int
	Col        int
	ID         string
	Op         string
	Expression interfaces.Expression
}

func NewUnarySum(lin int, col int, id string, op string, val interfaces.Expression) UnarySum {
	return UnarySum{lin, col, id, op, val}
}

func (p UnarySum) Execute(ast *environment.AST, env interface{}) interface{} {
	errorVal := environment.Symbol{Lin: 0, Col: 0, Type: environment.NULL, Value: nil}
	if env.(environment.Environment).VariableExists(p.ID) {
		foundVar := env.(environment.Environment).FindVar(p.ID)
		value := p.Expression.Execute(ast, env)

		if foundVar.Const {
			ast.SetError(p.Lin, p.Col, "No se puede actualizar el valor de una constante")
			foundVar.Value = nil
			env.(environment.Environment).UpdateVariable(p.ID, foundVar)
			return nil
		}

		if foundVar != errorVal {
			if p.Op == "+=" {
				if foundVar.Type == environment.FLOAT && (value.Type == environment.FLOAT || value.Type == environment.INTEGER) {
					switch v := value.Value.(type) {
					case int:
						foundVar.Value = foundVar.Value.(float64) + float64(v)
					case float64:
						foundVar.Value = foundVar.Value.(float64) + v

					}
					env.(environment.Environment).UpdateVariable(p.ID, foundVar)
					return nil
				} else if foundVar.Type == environment.INTEGER && (value.Type == environment.INTEGER) {
					switch v := value.Value.(type) {
					case int:
						foundVar.Value = foundVar.Value.(int) + v
					}
					env.(environment.Environment).UpdateVariable(p.ID, foundVar)
					return nil
				} else if foundVar.Type == environment.STRING && (value.Type == environment.STRING) {
					switch v := value.Value.(type) {
					case string:
						foundVar.Value = foundVar.Value.(string) + v
					}
					env.(environment.Environment).UpdateVariable(p.ID, foundVar)
					return nil
				} else {
					ast.SetError(p.Lin, p.Col, "No se puede sumar valor: " + value.GetType() + " a variable " + foundVar.GetType())
					foundVar.Value = nil
					env.(environment.Environment).UpdateVariable(p.ID, foundVar)
					return nil

				}
			} else if p.Op == "-=" {
				if foundVar.Type == environment.FLOAT && (value.Type == environment.FLOAT || value.Type == environment.INTEGER) {
					switch v := value.Value.(type) {
					case int:
						foundVar.Value = foundVar.Value.(float64) - float64(v)
					case float64:
						foundVar.Value = foundVar.Value.(float64) - v

					}
					env.(environment.Environment).UpdateVariable(p.ID, foundVar)
					return nil
				} else if foundVar.Type == environment.INTEGER && (value.Type == environment.INTEGER) {
					switch v := value.Value.(type) {
					case int:
						foundVar.Value = foundVar.Value.(int) - v
					}
					env.(environment.Environment).UpdateVariable(p.ID, foundVar)
					return nil
				} else {
					ast.SetError(p.Lin, p.Col, "No se puede sumar valor: " + value.GetType() + " a variable " + foundVar.GetType())
					foundVar.Value = nil
					env.(environment.Environment).UpdateVariable(p.ID, foundVar)
					return nil

				}
			}
		}

	} else if env.(environment.Environment).ReferenceExists(p.ID) {
		foundVar := env.(environment.Environment).FindReference(p.ID)
		value := p.Expression.Execute(ast, env)

		if foundVar.Const {
			ast.SetError(p.Lin, p.Col, "No se puede actualizar el valor de una constante")
			foundVar.Value = nil
			env.(environment.Environment).UpdateReference(p.ID, foundVar)
			return nil
		}

		if foundVar != errorVal {
			if p.Op == "+=" {
				if foundVar.Type == environment.FLOAT && (value.Type == environment.FLOAT || value.Type == environment.INTEGER) {
					switch v := value.Value.(type) {
					case int:
						foundVar.Value = foundVar.Value.(float64) + float64(v)
					case float64:
						foundVar.Value = foundVar.Value.(float64) + v

					}
					env.(environment.Environment).UpdateReference(p.ID, foundVar)
					return nil
				} else if foundVar.Type == environment.INTEGER && (value.Type == environment.INTEGER) {
					switch v := value.Value.(type) {
					case int:
						foundVar.Value = foundVar.Value.(int) + v
					}
					env.(environment.Environment).UpdateReference(p.ID, foundVar)
					return nil
				} else if foundVar.Type == environment.STRING && (value.Type == environment.STRING) {
					switch v := value.Value.(type) {
					case string:
						foundVar.Value = foundVar.Value.(string) + v
					}
					env.(environment.Environment).UpdateReference(p.ID, foundVar)
					return nil
				} else {

					ast.SetError(p.Lin, p.Col, "No se puede sumar valor: " + value.GetType() + " a variable " + foundVar.GetType())
					foundVar.Value = nil
					env.(environment.Environment).UpdateReference(p.ID, foundVar)
					return nil

				}
			} else if p.Op == "-=" {
				if foundVar.Type == environment.FLOAT && (value.Type == environment.FLOAT || value.Type == environment.INTEGER) {
					switch v := value.Value.(type) {
					case int:
						foundVar.Value = foundVar.Value.(float64) - float64(v)
					case float64:
						foundVar.Value = foundVar.Value.(float64) - v

					}
					env.(environment.Environment).UpdateReference(p.ID, foundVar)
					return nil
				} else if foundVar.Type == environment.INTEGER && (value.Type == environment.INTEGER) {
					switch v := value.Value.(type) {
					case int:
						foundVar.Value = foundVar.Value.(int) - v
					}
					env.(environment.Environment).UpdateReference(p.ID, foundVar)
					return nil
				} else {
					ast.SetError(p.Lin, p.Col, "No se puede restar valor: " + value.GetType() + " a variable " + foundVar.GetType())
					foundVar.Value = nil
					env.(environment.Environment).UpdateReference(p.ID, foundVar)
					return nil

				}
			}
		}
	}

	return nil
}
