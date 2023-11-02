package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type CallFuncExp struct {
	Lin        int
	Col        int
	Id         string
	Parameters []environment.FuncArg
}

func NewCallFuncExp(lin int, col int, val string, param []environment.FuncArg) CallFuncExp {
	exp := CallFuncExp{lin, col, val, param}
	return exp
}

func (p CallFuncExp) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	tmp1 := gen.NewTemp()
	size := env.(environment.Environment).Size["size"]
	gen.AddComment("LLAMANDO FUNCION")
	if len(p.Parameters) > 0 {
		gen.AddComment("-----guardando params-----")

		gen.AddExpression(tmp1, "P", strconv.Itoa(size+1), "+")
		for i := 0; i < len(p.Parameters); i++ {

			if strings.Contains(fmt.Sprintf("%T", p.Parameters[i].Value), "instructions") {
				result = p.Parameters[i].Value.(interfaces.Instruction).Execute(ast, env, gen)
			} else if strings.Contains(fmt.Sprintf("%T", p.Parameters[i].Value), "expressions") {
				result = p.Parameters[i].Value.(interfaces.Expression).Execute(ast, env, gen)
			} else {
				fmt.Println("Error en bloque")
			}

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

		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(p.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")

	}
	gen.AddComment("-----ENDCALL------")
	result.Value = tmp1

	return result
}
