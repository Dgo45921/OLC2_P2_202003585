package instructions

import (
	"PY1/environment"
	"PY1/generator"
)

type StructFuncCall struct {
	Lin    int
	Col    int
	Id     string
	FuncId string
}

func NewStructFuncCall(lin int, col int, id string, funcid string) StructFuncCall {
	return StructFuncCall{lin, col, id, funcid}
}

func (p StructFuncCall) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
func GetValueByArray(arr []string, symbol environment.Symbol) interface{} {
	var currentSymbol = symbol

	for _, key := range arr {
		found := false
		if kvArr, ok := currentSymbol.Value.([]environment.KeyValue); ok {
			for _, kv := range kvArr {
				if kv.Key == key {
					if _, isBreak := kv.Value.(environment.Symbol); isBreak {
						currentSymbol = kv.Value.(environment.Symbol)
					} else if _, isBreak := kv.Value.(environment.FunctionSymbol); isBreak {
						return kv.Value.(environment.FunctionSymbol)
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

	return currentSymbol
}
