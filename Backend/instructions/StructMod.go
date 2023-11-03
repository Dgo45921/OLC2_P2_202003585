package instructions

import (
	"PY1/environment"
	"PY1/generator"
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

func (p StructMod) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	foundVar := env.(environment.Environment).FindVar(p.ID)
	newvalue := p.Exp.Execute(ast, env, gen)
	if foundVar.Type == environment.STRUCT_IMP {

		err := UpdateValueByArray(p.Accesses, foundVar, newvalue, env)

		if err != nil {
			ast.SetError(p.Lin, p.Col, "no se pudo modificar el atributo")
		}

		return result

	} else {
		return result

	}
}

func UpdateValueByArray(arr []string, symbol environment.Symbol, val environment.Value, env interface{}) error {
	var currentSymbol = symbol

	// Create a reflect.Value of the provided value
	newVal := reflect.ValueOf(val)

	for _, key := range arr {
		found := false
		if kvArr, ok := currentSymbol.Value.([]environment.KeyValue); ok {
			for i, kv := range kvArr {
				if kv.Key == key {
					// Update kv.Value with the new value using reflection
					if _, isBreak := kv.Value.(environment.Value); isBreak {
						if !kv.Value.(environment.Value).Const && kv.Value.(environment.Value).Type == val.Type {
							if val.Type == environment.STRUCT_IMP {
								reflect.ValueOf(&kvArr[i]).Elem().FieldByName("Value").Set(newVal)
								found = true

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

func GetValueByArray2(arr []string, symbol environment.Symbol) interface{} {
	var currentSymbol = symbol
	var currentValue = environment.Value{
		Value:        "",
		IsTemp:       false,
		Type:         0,
		TrueLabel:    nil,
		FalseLabel:   nil,
		OutLabel:     nil,
		IntValue:     0,
		FloatValue:   0,
		BreakFlag:    false,
		ContinueFlag: false,
		ReturnFlag:   false,
		Dimentions:   nil,
		Const:        false,
		Scope:        0,
		Lin:          0,
		Col:          0,
		Id:           "",
		StructValues: nil,
	}

	for _, key := range arr {
		found := false
		if kvArr, ok := currentSymbol.Value.([]environment.KeyValue); ok {
			for _, kv := range kvArr {
				if kv.Key == key {
					currentValue = kv.Value.(environment.Value)
					if currentValue.Type == environment.STRUCT_IMP {
						currentSymbol.Value = currentValue.StructValues
					}

					found = true
					break
				}
			}
		}

		if !found {
			return nil
		}
	}

	return currentValue
}
