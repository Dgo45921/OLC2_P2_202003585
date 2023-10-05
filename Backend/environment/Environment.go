package environment

import (
	"fmt"
	"unsafe"
)

type Environment struct {
	Prev           interface{}
	SymbolTable    map[string]Symbol
	FunctionTable  map[string]FunctionSymbol
	ReferenceTable map[string]Symbol
	Scope          EnvType
}

func NewEnvironment(prev interface{}, scope EnvType) Environment {
	return Environment{
		Prev:           prev,
		SymbolTable:    make(map[string]Symbol),
		ReferenceTable: make(map[string]Symbol),
		FunctionTable:  make(map[string]FunctionSymbol),
		Scope:          scope,
	}
}

func (env Environment) PrintSymbolTableAddresses() {
	fmt.Println("Memory addresses of symbols in SymbolTable:")
	for key, symbol := range env.SymbolTable {
		address := unsafe.Pointer(&symbol)
		fmt.Printf("Key: %s, Address: %p\n", key, address)
	}
}

func (env Environment) VariableExists(id string) bool {
	var envTemporal = env
	for {
		if _, ok := envTemporal.SymbolTable[id]; ok {
			return true
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		return false

	}
}
func (env Environment) SaveVariable(id string, value Symbol) {
	env.SymbolTable[id] = value
}

func (env Environment) SaveReference(id string, value Symbol) {
	env.ReferenceTable[id] = value

}

func (env Environment) GetSymbolAddressByID(id string) unsafe.Pointer {
	symbol, exists := env.SymbolTable[id]
	if !exists {
		return nil
	}
	address := unsafe.Pointer(&symbol)
	return address
}

func (env Environment) FindReference(id string) Symbol {
	var envTemporal = env
	for {
		if foundVar, ok := envTemporal.ReferenceTable[id]; ok {
			return foundVar
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		return Symbol{Lin: 0, Col: 0, Type: NULL, Value: nil}

	}

}

func (env Environment) ReferenceExists(id string) bool {
	var envTemporal = env
	for {
		if _, ok := envTemporal.ReferenceTable[id]; ok {
			return true
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		return false

	}
}

func (env Environment) FuncExists(id string) bool {

	if _, ok := env.FunctionTable[id]; ok {
		return true
	}

	return false
}
func (env Environment) SaveFunc(id string, value FunctionSymbol) {
	env.FunctionTable[id] = value
}

func (env Environment) SaveStruct(id string, value Symbol) {
	env.SymbolTable[id] = value
}

func (env Environment) UpdateVariable(id string, value Symbol) {
	var envTemporal = env
	for {
		if _, ok := envTemporal.SymbolTable[id]; ok {
			envTemporal.SymbolTable[id] = value
			break
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		break

	}

}

func (env Environment) SetReferenceValues(realid string, secondaryID string) {
	result := env.UpdateValue(realid, env.ReferenceTable[secondaryID])
	if result == false {
		env.UpdateValueReference(realid, env.ReferenceTable[secondaryID])
	}
}

func (env Environment) UpdateValue(id string, value Symbol) bool {
	var envTemporal = env
	for {
		for key, _ := range envTemporal.SymbolTable {
			if key == id {
				val := value
				pivote := envTemporal.SymbolTable[key]
				pivote.Value = val.Value

				envTemporal.SymbolTable[key] = pivote
				return true
			}
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		break

	}
	return false
}

func (env Environment) UpdateValueReference(id string, value Symbol) {
	var envTemporal = env
	for {
		for key, _ := range envTemporal.ReferenceTable {
			if key == id {
				val := value
				pivote := envTemporal.ReferenceTable[key]
				pivote.Value = val.Value

				envTemporal.ReferenceTable[key] = pivote
				return
			}
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		break

	}

}

func (env Environment) UpdateReference(id string, value Symbol) {
	var envTemporal = env
	for {
		if _, ok := envTemporal.ReferenceTable[id]; ok {
			envTemporal.ReferenceTable[id] = value
			break
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		break

	}

}

func (env Environment) FindVar(id string) Symbol {
	var envTemporal = env
	for {
		if foundVar, ok := envTemporal.SymbolTable[id]; ok {
			return foundVar
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		return Symbol{Lin: 0, Col: 0, Type: NULL, Value: nil}

	}

}

func (env Environment) FindFunc(id string) (FunctionSymbol, bool) {
	var envTemporal = env
	for {
		if foundVar, ok := envTemporal.FunctionTable[id]; ok {
			return foundVar, true
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		return FunctionSymbol{Lin: 0, Col: 0}, false

	}

}

func (env Environment) InsideLoop() bool {
	var tmpEnv = env
	for {
		if tmpEnv.Scope == WHILE || tmpEnv.Scope == FOR || tmpEnv.Scope == CASE || tmpEnv.Scope == DEFAULT {
			return true
		}
		if tmpEnv.Prev == nil {
			break
		} else {
			tmpEnv = tmpEnv.Prev.(Environment)
		}
	}
	return false
}
