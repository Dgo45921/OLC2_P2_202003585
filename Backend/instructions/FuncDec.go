package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
	"strings"
)

type FuncDec struct {
	Lin        int
	Col        int
	Id         string
	Args       []environment.FuncParam
	ReturnType interface{}
	insBlock   []interface{}
	Mutating   bool
}

func NewFuncDec(lin int, col int, id string, args []environment.FuncParam, ret interface{}, insb []interface{}, mut bool) FuncDec {
	return FuncDec{lin, col, id, args, ret, insb, mut}
}

func (p FuncDec) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	gen.SetMainFlag(false)
	gen.AddComment("******** Funcion " + p.Id + " ********")
	gen.AddTittle(p.Id)
	//entorno
	var envFunc environment.Environment
	envFunc = environment.NewEnvironment(env.(environment.Environment), environment.FUNC)
	envFunc.Size["size"] = envFunc.Size["size"] + 1
	//variables
	for _, s := range p.Args {
		res := prueba(s.SID, s.Type)
		envFunc.SaveVariable(res.Value, res.Type)
	}
	//instrucciones func
	for _, s := range p.insBlock {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			s.(interfaces.Instruction).Execute(ast, envFunc, gen)

			////agregando etiquetas de salida
			//for _, lvl := range resInst.OutLabel {
			//	gen.AddLabel(lvl.(string))
			//}

		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			result = s.(interfaces.Expression).Execute(ast, envFunc, gen)
			//agregando etiquetas de salida
			for _, lvl := range result.OutLabel {
				gen.AddLabel(lvl.(string))
			}
		} else {
			fmt.Println("Error en bloque")
		}
	}
	gen.AddEnd()
	gen.SetMainFlag(true)

	if _, isBreak := p.ReturnType.(string); isBreak {
		result.Type = getReturnType(p.ReturnType.(string))
	} else {
		result.Type = environment.NULL
	}

	return result
}
func getReturnType(str string) environment.TipoExpresion {
	if str == "String" {
		return environment.STRING
	} else if str == "Int" {
		return environment.INTEGER
	} else if str == "Float" {
		return environment.FLOAT
	} else if str == "Bool" {
		return environment.BOOLEAN
	} else if str == "Character" {
		return environment.CHAR
	} else {
		return environment.STRUCT_IMP
	}
}

func prueba(id string, tyype string) environment.Value {

	var result environment.Value
	result = environment.NewValue(id, false, getReturnType(tyype))

	return result
}
