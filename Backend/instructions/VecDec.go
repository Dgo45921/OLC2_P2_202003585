package instructions

import (
	"PY1/environment"
	"PY1/expressions"
	"PY1/generator"
	"PY1/interfaces"
	"strconv"
)

type VecDec struct {
	Lin     int
	Col     int
	Id      string
	Type    string
	DefType interface{}
	Exp     interface{}
}

func NewVecDec(lin int, col int, id string, tyype string, deftype interface{}, exp interface{}) VecDec {
	NewVecDeclaration := VecDec{lin, col, id, tyype, deftype, exp}
	return NewVecDeclaration
}

func (p VecDec) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result, val environment.Value
	size := len(p.Exp.(expressions.Vector).Value)
	if env.(environment.Environment).VariableExists(p.Id) {
		ast.SetError(p.Lin, p.Col, "Error, variable ya declarada!")
		return environment.Value{}
	}
	//generando array
	gen.AddComment("Generando array")
	newtmp1 := gen.NewTemp()
	newtmp2 := gen.NewTemp()
	gen.AddAssign(newtmp1, "H")
	gen.AddExpression(newtmp2, newtmp1, "1", "+")
	gen.AddSetHeap("(int)H", strconv.Itoa(size))
	gen.AddExpression("H", "H", strconv.Itoa(size+1), "+")
	//recorriendo lista de expressiones
	for _, s := range p.Exp.(expressions.Vector).Value {
		val = s.(interfaces.Expression).Execute(ast, env, gen)
		gen.AddSetHeap("(int)"+newtmp2, val.Value)
		gen.AddExpression(newtmp2, newtmp2, "1", "+")
	}
	result = environment.Value{
		Value:        newtmp1,
		IsTemp:       true,
		Type:         getType(val.Type),
		TrueLabel:    nil,
		FalseLabel:   nil,
		OutLabel:     nil,
		IntValue:     0,
		FloatValue:   0,
		BreakFlag:    false,
		ContinueFlag: false,
	}

	newVar:= env.(environment.Environment).SaveVector(p.Id, result.Type, size)

	gen.AddSetStack(strconv.Itoa(newVar.Position), result.Value)
	gen.AddBr()

	return result

}

func getType(val environment.TipoExpresion) environment.TipoExpresion {
	if val == environment.INTEGER {
		return environment.VECTOR_INT
	} else if val == environment.FLOAT {
		return environment.VECTOR_FLOAT
	} else if val == environment.BOOLEAN {
		return environment.VECTOR_BOOLEAN
	} else if val == environment.CHAR {
		return environment.VECTOR_CHAR
	} else if val == environment.STRING {
		return environment.VECTOR_STRING
	}

	return environment.VECTOR
}
