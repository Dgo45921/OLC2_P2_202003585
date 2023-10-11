package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"reflect"
)

type VectorModification struct {
	Lin        int
	Col        int
	DestinyID  string
	Indexes    []interface{}
	Expression interfaces.Expression
}

func NewVectorModification(lin int, col int, id string, indexes []interface{}, val interfaces.Expression) VectorModification {
	asig := VectorModification{lin, col, id, indexes, val}
	return asig
}

func (p VectorModification) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}

func setIndexValue(arr interface{}, newValue interface{}, indexes []int) bool {
	if len(indexes) == 0 || reflect.ValueOf(arr).Kind() != reflect.Slice {
		return false
	}

	index := indexes[0]
	if index < 0 || index >= reflect.ValueOf(arr).Len() {
		return false
	}

	if len(indexes) == 1 {
		reflect.ValueOf(arr).Index(index).Set(reflect.ValueOf(newValue))
		return true
	}

	nextLevel := reflect.ValueOf(arr).Index(index).Interface()
	return setIndexValue(nextLevel, newValue, indexes[1:])
}



func AllNonNegativeIntegers(arr []int) bool {
	for _, num := range arr {
		if num < 0 {
			return false
		}
	}
	return true
}

func GetIndexValue(arr interface{}, indexes []int) (interface{}, bool) {
	if len(indexes) == 0 || reflect.ValueOf(arr).Kind() != reflect.Slice {
		return nil, false
	}

	index := indexes[0]
	if index < 0 || index >= reflect.ValueOf(arr).Len() {
		return nil, false
	}

	if len(indexes) == 1 {
		return reflect.ValueOf(arr).Index(index).Interface(), true
	}

	nextLevel := reflect.ValueOf(arr).Index(index).Interface()
	return GetIndexValue(nextLevel, indexes[1:])
}
