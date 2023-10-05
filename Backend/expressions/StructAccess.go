package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
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

func (p StructAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	if env.(environment.Environment).VariableExists(p.ID) {
		foundVar := env.(environment.Environment).FindVar(p.ID)
		if foundVar.Type == environment.STRUCT_IMP {
			foundSymbol := GetValueByArray(p.Accesses, foundVar)
			if foundSymbol != nil {
				if _, isBreak := foundSymbol.(interfaces.Expression); isBreak {
					foundSymbol = foundSymbol.(interfaces.Expression).Execute(ast, env)
					return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: foundSymbol.(environment.Symbol).Type, Value: foundSymbol.(environment.Symbol).Value}
				}
				return foundSymbol.(environment.Symbol)
			}
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}

		} else {
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}

		}
	} else if env.(environment.Environment).ReferenceExists(p.ID) {
		foundVar := env.(environment.Environment).FindReference(p.ID)
		if foundVar.Type == environment.STRUCT_IMP {
			foundSymbol := GetValueByArray(p.Accesses, foundVar)
			if foundSymbol != nil {
				if _, isBreak := foundSymbol.(interfaces.Expression); isBreak {
					foundSymbol = foundSymbol.(interfaces.Expression).Execute(ast, env)
					return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: foundSymbol.(environment.Symbol).Type, Value: foundSymbol.(environment.Symbol).Value}
				}
				return foundSymbol.(environment.Symbol)
			}
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}

		} else {
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}

		}

	}

	return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
}

func GetValueByArray(arr []string, symbol environment.Symbol) interface{} {
	var currentSymbol = symbol

	for _, key := range arr {
		found := false
		if kvArr, ok := currentSymbol.Value.([]environment.KeyValue); ok {
			for _, kv := range kvArr {
				if kv.Key == key {
					currentSymbol = kv.Value.(environment.Symbol)
					found = true
					break
				}
			}
		}

		if !found {
			return nil
		}
	}

	return currentSymbol
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
