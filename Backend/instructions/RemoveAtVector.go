package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type RemoveAtVector struct {
	Lin   int
	Col   int
	Id    string
	Index interfaces.Expression
}

func NewRemoveAtVector(lin int, col int, id string, index interfaces.Expression) RemoveAtVector {
	asig := RemoveAtVector{lin, col, id, index}
	return asig
}

func (p RemoveAtVector) Execute(ast *environment.AST, env interface{}) interface{} {
	if env.(environment.Environment).VariableExists(p.Id) {
		foundVar := env.(environment.Environment).FindVar(p.Id)
		if foundVar.Const {
			ast.SetError(p.Lin, p.Col, "No se puede modificar un vector constante")
			return nil

		}
		indexval := p.Index.Execute(ast, env)
		if slice, isArray := foundVar.Value.([]interface{}); isArray {
			if indexval.Type == environment.INTEGER {
				if indexval.Value.(int) >= 0 && indexval.Value.(int) < len(slice) && !foundVar.Const {
					foundVar.Value = append(slice[:indexval.Value.(int)], slice[indexval.Value.(int)+1:]...)
					env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				} else {
					ast.SetError(p.Lin, p.Col, "Indice no disponible en el vector")
				}
			} else {
				ast.SetError(p.Lin, p.Col, "el indice luego de 'at:' debe ser un entero")
			}
		} else {
			ast.SetError(p.Lin, p.Col, "la funcion remove(at) solo funciona en vectores")
		}
		return nil

	} else if env.(environment.Environment).ReferenceExists(p.Id) {
		foundVar := env.(environment.Environment).FindReference(p.Id)
		if foundVar.Const {
			ast.SetError(p.Lin, p.Col, "No se puede modificar un vector constante")
			return nil

		}
		indexval := p.Index.Execute(ast, env)
		if slice, isArray := foundVar.Value.([]interface{}); isArray {
			if indexval.Type == environment.INTEGER {
				if indexval.Value.(int) >= 0 && indexval.Value.(int) < len(slice) && !foundVar.Const {
					foundVar.Value = append(slice[:indexval.Value.(int)], slice[indexval.Value.(int)+1:]...)
					env.(environment.Environment).UpdateReference(p.Id, foundVar)
				} else {
					ast.SetError(p.Lin, p.Col, "Indice no disponible en el vector")
				}
			} else {
				ast.SetError(p.Lin, p.Col, "el indice luego de 'at:' debe ser un entero")
			}
		} else {
			ast.SetError(p.Lin, p.Col, "la funcion remove(at) solo funciona en vectores")
		}
		return nil

	}
	return nil
}
