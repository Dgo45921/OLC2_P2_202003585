package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"reflect"
)

type VectorAccess struct {
	Lin   int
	Col   int
	Id    string
	Index []interface{}
}

func NewVectorAccess(lin int, col int, id string, index []interface{}) VectorAccess {
	exp := VectorAccess{lin, col, id, index}
	return exp
}

func (p VectorAccess) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
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
