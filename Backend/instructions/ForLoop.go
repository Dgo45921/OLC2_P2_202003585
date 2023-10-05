package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type For struct {
	Lin      int
	Col      int
	Id       string
	Range    interfaces.Expression
	insBlock []interface{}
}

func NewFor(lin int, col int, id string, rangge interfaces.Expression, insBlock []interface{}) For {
	return For{lin, col, id, rangge, insBlock}
}

func (p For) Execute(ast *environment.AST, env interface{}) interface{} {

	var iterable = p.Range.Execute(ast, env)
	if iterable.Type == environment.VECTOR {
		return nil
	}

	if iterable.Type == environment.VECTOR_BOOLEAN || iterable.Type == environment.VECTOR_INT || iterable.Type == environment.VECTOR_STRING || iterable.Type == environment.VECTOR_FLOAT || iterable.Type == environment.VECTOR_CHAR || iterable.Type == environment.VECTOR_STRUCT {
		var iterableType = environment.INTEGER
		switch iterable.Type {
		case environment.VECTOR_BOOLEAN:
			iterableType = environment.BOOLEAN
		case environment.VECTOR_INT:
			iterableType = environment.INTEGER
		case environment.VECTOR_STRING:
			iterableType = environment.STRING
		case environment.VECTOR_FLOAT:
			iterableType = environment.FLOAT
		case environment.VECTOR_CHAR:
			iterableType = environment.CHAR
		case environment.VECTOR_STRUCT:
			iterableType = environment.STRUCT_IMP

		}

		arr := iterable.Value.([]interface{})
		arrLen := len(arr)
	outerloop:
		for index := 0; index < arrLen; index++ {
			element := arr[index]

			var currentIterable = environment.Symbol{
				Value: element,
				Type:  iterableType,
				Scope: env.(environment.Environment).Scope,
			}
			var newEnv = environment.NewEnvironment(env, environment.FOR)
			newEnv.SaveVariable(p.Id, currentIterable)
			ast.SaveSymbol(p.Id,currentIterable)

			for _, inst := range p.insBlock {

				if as, isasignation := inst.(Asignation); isasignation {
					if as.Id == p.Id {
						ast.SetError(p.Lin, p.Col, "no se puede cambiar el valor del iterable")
						return nil
					}
				}

				var response = inst.(interfaces.Instruction).Execute(ast, newEnv)
				if response != nil {
					if _, isBreak := response.(Break); isBreak {
						return nil
					} else if _, isContinue := response.(Continue); isContinue {
						continue outerloop
					} else if _, isReturn := response.(environment.Symbol); isReturn {
						return response

					}
				}

			}
		}
	} else if iterable.Type == environment.STRING {
		var iterableType = environment.CHAR
		arr := iterable.Value.(string)
		arrLen := len(arr)
	outerloop2:
		for index := 0; index < arrLen; index++ {
			element := arr[index]

			var currentIterable = environment.Symbol{
				Value: element,
				Type:  iterableType,
				Scope: env.(environment.Environment).Scope,
			}
			var newEnv = environment.NewEnvironment(env, environment.FOR)
			newEnv.SaveVariable(p.Id, currentIterable)
			ast.SaveSymbol(p.Id,currentIterable)

			for _, inst := range p.insBlock {

				if as, isasignation := inst.(Asignation); isasignation {
					if as.Id == p.Id {
						ast.SetError(p.Lin, p.Col, "no se puede cambiar el valor del iterable")
						return nil
					}
				}

				var response = inst.(interfaces.Instruction).Execute(ast, newEnv)
				if response != nil {
					if _, isBreak := response.(Break); isBreak {
						return nil
					} else if _, isContinue := response.(Continue); isContinue {
						continue outerloop2
					}
				}

			}
		}
	} else {
		ast.SetError(p.Lin, p.Col, "no se recibio un iterable")
		return nil
	}

	return nil
}
