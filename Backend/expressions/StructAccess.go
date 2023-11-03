package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"fmt"
)

type StructAccess struct {
	Lin      int
	Col      int
	ID       string
	Accesses []string
}

func NewStructAccess(lin int, col int, id string, accesses []string) StructAccess {
	structaccess := StructAccess{lin, col, id, accesses}
	return structaccess
}

func (p StructAccess) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	if env.(environment.Environment).VariableExists(p.ID) {
		foundVar := env.(environment.Environment).FindVar(p.ID)
		if foundVar.Type == environment.STRUCT_IMP {
			foundSymbol := GetValueByArray(p.Accesses, foundVar)
			if foundSymbol != nil {
				if _, isBreak := foundSymbol.(environment.Value); isBreak {
					return foundSymbol.(environment.Value)
				}

			}
			return result

		} else {
			return result

		}
	}
	return result
}

func GetValueByArray(arr []string, symbol environment.Symbol) interface{} {
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

func searchNestedValue(data environment.KeyValue, keys []string) (interface{}, error) {
	stack := []environment.KeyValue{{Key: "", Value: data}}

	for _, key := range keys {
		var newStack []environment.KeyValue

		for _, kv := range stack {
			if nestedData, ok := kv.Value.(environment.KeyValue); ok && nestedData.Key == key {
				newStack = append(newStack, nestedData)
			}
		}

		if len(newStack) == 0 {
			return nil, fmt.Errorf("key '%s' not found", key)
		}

		stack = newStack
	}

	if len(stack) > 0 {
		return stack[len(stack)-1].Value, nil
	}

	return nil, fmt.Errorf("value not found")
}
