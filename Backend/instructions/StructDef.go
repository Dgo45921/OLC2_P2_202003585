package instructions

import (
	"PY1/environment"
)

type StructDef struct {
	Lin      int
	Col      int
	Id       string
	insBlock []interface{}
	Type     environment.TipoExpresion
}

func NewStructDef(lin int, col int, id string, insBlock []interface{}) StructDef {
	return StructDef{lin, col, id, insBlock, environment.STRUCT_DEF}
}

func (p StructDef) Execute(ast *environment.AST, env interface{}) interface{} {
	if env.(environment.Environment).Prev != nil {
		ast.SetError(p.Lin, p.Col, "Los structs solo pueden ser declarados en el ambito global")
		return nil
	}

	structMap := []environment.KeyValue{}

	for _, inst := range p.insBlock {

		if _, isVarDec := inst.(VarDec); isVarDec {
			response := inst.(VarDec).GetVarDec(ast, env)
			if response != nil {
				if !repeatedValue(inst.(VarDec).Id, structMap) {
					newKeyValuePair := environment.KeyValue{inst.(VarDec).Id, response}
					structMap = append(structMap, newKeyValuePair)
				} else {
					ast.SetError(p.Lin, p.Col, "atributo var en struct repetido")
					return nil

				}

			} else {
				ast.SetError(p.Lin, p.Col, "El tipo de asignacion a un atributo var no fue válida")
				return nil

			}

		} else if _, isConstDec := inst.(ConstDec); isConstDec {

			response := inst.(ConstDec).GetConstDec(ast, env)
			if response != nil {
				if !repeatedValue(inst.(ConstDec).Id, structMap) {
					newKeyValuePair := environment.KeyValue{inst.(ConstDec).Id, response}
					structMap = append(structMap, newKeyValuePair)
				} else {
					ast.SetError(p.Lin, p.Col, "atributo const en struct repetido")
					return nil
				}
			} else {
				ast.SetError(p.Lin, p.Col, "El tipo de asignacion a un atributo const no fue válida")
				return nil

			}

		} else if _, isVecDec := inst.(VecDec); isVecDec {
			response := inst.(VecDec).GetVecDec(ast, env)
			if response != nil {
				if !repeatedValue(inst.(VecDec).Id, structMap) {
					newKeyValuePair := environment.KeyValue{inst.(VecDec).Id, response}
					structMap = append(structMap, newKeyValuePair)
				} else {
					ast.SetError(p.Lin, p.Col, "atributo const en struct repetido")
					return nil

				}
			} else {
				ast.SetError(p.Lin, p.Col, "El tipo de asignacion a un atributo vector no fue válida")
				return nil

			}

		} else if _, isMatrixDec := inst.(MatrixDec); isMatrixDec {

			response := inst.(MatrixDec).GetMatrixDec(ast, env)
			if response != nil {
				if !repeatedValue(inst.(MatrixDec).Id, structMap) {
					newKeyValuePair := environment.KeyValue{inst.(MatrixDec).Id, response}
					structMap = append(structMap, newKeyValuePair)
				} else {
					ast.SetError(p.Lin, p.Col, "atributo const en struct repetido")
					return nil

				}
			} else {
				ast.SetError(p.Lin, p.Col, "Tipo de asignacion a matriz no fue valida")
				return nil

			}

		}else if _, isFuncDec := inst.(FuncDec); isFuncDec {

			response := inst.(FuncDec).GetFuncDec(ast, env)
			if response != nil {
				if !repeatedValue(inst.(FuncDec).Id, structMap) {
					newKeyValuePair := environment.KeyValue{inst.(FuncDec).Id, response}
					structMap = append(structMap, newKeyValuePair)
				} else {
					ast.SetError(p.Lin, p.Col, "atributo const en struct repetido")
					return nil

				}
			} else {
				ast.SetError(p.Lin, p.Col, "Tipo de asignacion a matriz no fue valida")
				return nil

			}

		}

	}

	newstruct := environment.Symbol{
		Lin:        p.Lin,
		Col:        p.Col,
		Value:      structMap,
		Type:       p.Type,
		StructType: p.Id,
	}

	env.(environment.Environment).SaveStruct(p.Id, newstruct)

	return nil
}

func repeatedValue(id string, arraykevalue []environment.KeyValue) bool {
	for _, kv := range arraykevalue {
		if id == kv.Key {
			return true
		}
	}

	return false
}
