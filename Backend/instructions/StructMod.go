package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
	"errors"

	"reflect"
)

type StructMod struct {
	Lin      int
	Col      int
	ID       string
	Accesses []string
	Exp      interfaces.Expression
}

func NewStructMod(lin int, col int, id string, accesses []string, exp interfaces.Expression) StructMod {
	structaccess := StructMod{lin, col, id, accesses, exp}
	return structaccess
}

func (p StructMod) Execute(ast *environment.AST, env interface{}) interface{} {
	foundVar := env.(environment.Environment).FindVar(p.ID)
	newvalue := p.Exp.Execute(ast, env)
	if foundVar.Type == environment.STRUCT_IMP {

		err := UpdateValueByArray(p.Accesses, foundVar, newvalue, env)

		if err != nil {
			ast.SetError(p.Lin, p.Col, "no se pudo modificar el atributo")
		}

		return nil

	} else {
		return nil

	}

}

func UpdateValueByArray(arr []string, symbol environment.Symbol, val environment.Symbol, env interface{}) error {
	var currentSymbol = symbol

	// Create a reflect.Value of the provided value
	newVal := reflect.ValueOf(val)

	for _, key := range arr {
		found := false
		if kvArr, ok := currentSymbol.Value.([]environment.KeyValue); ok {
			for i, kv := range kvArr {
				if kv.Key == key {
					// Update kv.Value with the new value using reflection
					if _, isBreak := kv.Value.(environment.Symbol); isBreak {
						if !kv.Value.(environment.Symbol).Const && kv.Value.(environment.Symbol).Type == val.Type {
							if val.Type == environment.STRUCT_IMP {
								foundVar := env.(environment.Environment).FindVar(val.StructType)
								if foundVar.StructType == val.StructType {
									reflect.ValueOf(&kvArr[i]).Elem().FieldByName("Value").Set(newVal)
									found = true
								} else {
									return errors.New("Struct type mismatch")
								}

							} else {
								reflect.ValueOf(&kvArr[i]).Elem().FieldByName("Value").Set(newVal)
								found = true
							}
						} else {
							return errors.New("Invalid symbol type or constant")
						}
					}
				}
			}
		}

		if !found {
			return errors.New("Key not found")
		}
	}

	return nil
}
