package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type CallFuncInst struct {
	Lin        int
	Col        int
	Id         string
	Parameters []environment.FuncArg
}

func NewCallFuncInst(lin int, col int, val string, param []environment.FuncArg) CallFuncInst {
	exp := CallFuncInst{lin, col, val, param}
	return exp
}

func (p CallFuncInst) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	size := env.(environment.Environment).Size["size"]
	gen.AddComment("Inicio de llamada")
	if len(p.Parameters) > 0 {
		tmp1 := gen.NewTemp()
		gen.AddExpression(tmp1, "P", strconv.Itoa(size+1), "+")
		for i := 0; i < len(p.Parameters); i++ {

			//ejecutando parametros
			if strings.Contains(fmt.Sprintf("%T", p.Parameters[i].Value), "instructions") {
				result = p.Parameters[i].Value.(interfaces.Instruction).Execute(ast, env, gen)
			} else if strings.Contains(fmt.Sprintf("%T", p.Parameters[i].Value), "expressions") {
				result = p.Parameters[i].Value.(interfaces.Expression).Execute(ast, env, gen)
			} else {
				fmt.Println("Error en bloque")
			}
			//guardando
			gen.AddSetStack("(int)"+tmp1, result.Value)
			if (len(p.Parameters) - 1) != i {
				gen.AddExpression(tmp1, tmp1, "1", "+")
			}
		}

		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(p.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")

	} else {
		tmp1 := gen.NewTemp()

		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(p.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")

	}
	gen.AddComment("Final de llamada")
	return result
}
