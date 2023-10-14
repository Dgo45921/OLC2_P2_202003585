package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
)

type ArithmeticOperation struct {
	Lin      int
	Col      int
	OpIzq    interfaces.Expression
	Operator string
	OpDer    interfaces.Expression
}

func NewArithmeticOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) ArithmeticOperation {
	exp := ArithmeticOperation{Lin: lin, Col: col, OpIzq: Op1, Operator: Operador, OpDer: Op2}
	return exp
}

func (o ArithmeticOperation) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var dominante environment.TipoExpresion

	tablaDominante := [5][5]environment.TipoExpresion{
		//		INTEGER			FLOAT				STRING				BOOLEAN				NULL
		{environment.INTEGER, environment.FLOAT, environment.STRING, environment.BOOLEAN, environment.NULL},
		//FLOAT
		{environment.FLOAT, environment.FLOAT, environment.STRING, environment.NULL, environment.NULL},
		//STRING
		{environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.NULL},
		//BOOLEAN
		{environment.BOOLEAN, environment.NULL, environment.STRING, environment.BOOLEAN, environment.NULL},
		//NULL
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}

	var op1, op2, result environment.Value

	newTemp := gen.NewTemp()
	switch o.Operator {
	case "+":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			op2 = o.OpDer.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}

			//validar tipo dominante
			dominante = tablaDominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "+")
				result = environment.NewValue(newTemp, true, dominante)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.STRING {
				//llamar a generar concatstring

				//gen.GenerateConcatString()

				{

					gen.AddComment("Concat Strigns")
					tmp := gen.NewTemp()
					gen.AddAssign(tmp, "H")
					gen.AddComment("Op izq")
					//NO AGREGO -1 PORQUE QUIERO CONCATENAR
					labelRepetir := gen.NewLabel()
					labelSalir := gen.NewLabel()
					//Temporal del string
					tmpInicio := gen.NewTemp()
					gen.AddAssign(tmpInicio, op1.Value)
					gen.AddLabel(labelRepetir)
					tmpIterar := gen.NewTemp()
					gen.AddAssign(tmpIterar, "heap[(int)"+tmpInicio+"]")
					labelTrue := gen.NewLabel()
					gen.AddIf(tmpIterar, "-1", "!=", labelTrue)
					gen.AddGoto(labelSalir)
					gen.AddLabel(labelTrue)
					gen.AddSetHeap("(int)H", tmpIterar)
					gen.AddExpression("H", "H", "1", "+")
					gen.AddExpression(tmpInicio, tmpInicio, "1", "+")
					gen.AddGoto(labelRepetir)
					//encuentra en el heap -1
					gen.AddLabel(labelSalir)
					gen.AddComment("Op der")
					//NO AGREGO -1 PORQUE QUIERO CONCATENAR
					labelRepetir = gen.NewLabel()
					labelSalir = gen.NewLabel()
					//Temporal del string
					tmpInicio = gen.NewTemp()
					gen.AddAssign(tmpInicio, op2.Value)
					gen.AddLabel(labelRepetir)
					tmpIterar = gen.NewTemp()
					gen.AddAssign(tmpIterar, "heap[(int)"+tmpInicio+"]")
					labelTrue = gen.NewLabel()
					gen.AddIf(tmpIterar, "-1", "!=", labelTrue)
					gen.AddGoto(labelSalir)
					gen.AddLabel(labelTrue)
					gen.AddSetHeap("(int)H", tmpIterar)
					gen.AddExpression("H", "H", "1", "+")
					gen.AddExpression(tmpInicio, tmpInicio, "1", "+")
					gen.AddGoto(labelRepetir)
					//encuentra en el heap -1
					gen.AddLabel(labelSalir)

					gen.AddSetHeap("(int)H", "-1")
					gen.AddExpression("H", "H", "1", "+")
					gen.AddComment("SALIENDO")
					result = environment.NewValue(tmp, true, environment.STRING)
					return result

				}
			} else {
				fmt.Println("ERROR: No es posible sumar ", dominante)
				return result
			}
		}
	case "-":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			op2 = o.OpDer.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "-")
				result = environment.NewValue(newTemp, true, dominante)
				result.IntValue = op1.IntValue - op2.IntValue
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible restar")
			}
		}
	case "*":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			op2 = o.OpDer.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "*")
				result = environment.NewValue(newTemp, true, dominante)
				result.IntValue = op1.IntValue * op2.IntValue
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible multiplicar")
			}
		}
	case "/":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			op2 = o.OpDer.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				lvl1 := gen.NewLabel()
				lvl2 := gen.NewLabel()

				gen.AddIf(op2.Value, "0", "!=", lvl1)
				gen.AddPrintf("c", "77")
				gen.AddPrintf("c", "97")
				gen.AddPrintf("c", "116")
				gen.AddPrintf("c", "104")
				gen.AddPrintf("c", "69")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "111")
				gen.AddPrintf("c", "114")
				gen.AddExpression(newTemp, "0", "", "")
				gen.AddGoto(lvl2)
				gen.AddLabel(lvl1)
				gen.AddExpression(newTemp, op1.Value, op2.Value, "/")
				gen.AddLabel(lvl2)
				result = environment.NewValue(newTemp, true, dominante)
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible dividir")
			}

		}

	case "%":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			op2 = o.OpDer.Execute(ast, env, gen)
			if op1.Type < 0 || int(op1.Type) >= len(tablaDominante) || op2.Type < 0 || int(op2.Type) >= len(tablaDominante) {
				ast.SetError(o.Lin, o.Col, "Error, tipo de operacion no valida!")
				return environment.Value{}
			}
			dominante = tablaDominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				lvl1 := gen.NewLabel()
				lvl2 := gen.NewLabel()

				gen.AddIf(op2.Value, "0", "!=", lvl1)
				gen.AddPrintf("c", "77")
				gen.AddPrintf("c", "97")
				gen.AddPrintf("c", "116")
				gen.AddPrintf("c", "104")
				gen.AddPrintf("c", "69")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "111")
				gen.AddPrintf("c", "114")
				gen.AddExpression(newTemp, "0", "", "")
				gen.AddGoto(lvl2)
				gen.AddLabel(lvl1)
				//gen.AddExpression(newTemp, op1.Value, op2.Value, "%")
				gen.AddExpression(newTemp, "fmod("+op1.Value, op2.Value+")", ",")
				gen.AddLabel(lvl2)
				result = environment.NewValue(newTemp, true, dominante)
				return result
			} else {
				ast.SetError(o.Lin, o.Col, "ERROR: No es posible hacer modulo")
			}

		}

	case "Int":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			gen.AddComment("----INT CAST---")
			t2 := gen.NewTemp()
			t3 := gen.NewTemp()
			t4 := gen.NewTemp()
			tpeso := gen.NewTemp()

			L1 := gen.NewLabel()
			L2 := gen.NewLabel()
			L3 := gen.NewLabel()
			L4 := gen.NewLabel()
			L5 := gen.NewLabel()

			gen.AddAssign(tpeso, "1")
			gen.AddAssign(t2, op1.Value)
			gen.AddAssign(t3, "0")
			gen.AddLabel(L4)
			gen.AddAssign(t4, "heap[(int)"+t2+"]")
			gen.AddIf(t4, "-1", "==", L1)
			gen.AddComment("viendo si solo numeros")
			gen.AddIf(t4, "48", "<", L5)
			gen.AddIf(t4, "57", ">", L5)

			gen.AddGoto(L2)
			gen.AddLabel(L1)

			gen.AddGoto(L3)
			gen.AddLabel(L2)
			gen.AddAssign(t3, t2)
			gen.AddExpression(t2, t2, "1", "+")
			gen.AddGoto(L4)
			gen.AddLabel(L3)

			t5 := gen.NewTemp()
			t6 := gen.NewTemp()
			L6 := gen.NewLabel()
			L7 := gen.NewLabel()
			tresultado := gen.NewTemp()
			LSalir := gen.NewLabel()

			gen.AddAssign(t5, t3)
			LRec := gen.NewLabel()

			gen.AddLabel(LRec)
			gen.AddAssign(t6, "heap[(int)"+t5+"]")
			gen.AddIf(t6, "-2", "==", L6)
			gen.AddIf(t6, "-3", "==", L5)
			gen.AddIf(t6, "-1", "==", LSalir)
			gen.AddIf(t5, "-1", "==", LSalir)
			gen.AddGoto(L7)

			gen.AddLabel(L6)
			gen.AddExpression(t5, t5, "1", "-")
			gen.AddAssign(tresultado, "heap[(int)"+t5+"]")
			gen.AddGoto(LSalir)

			gen.AddLabel(L7)
			tmpComparar := gen.NewTemp()
			L8 := gen.NewLabel()
			L9 := gen.NewLabel()
			gen.AddAssign(tmpComparar, "heap[(int)"+t5+"]")
			gen.AddIf(tpeso, "1", "==", L8)
			gen.AddGoto(L9)
			gen.AddLabel(L8)

			tmpOperar := gen.NewTemp()
			gen.AddExpression(tmpComparar, tmpComparar, "48", "-")
			gen.AddExpression(tmpOperar, tmpComparar, "1", "*")
			gen.AddExpression(tresultado, tresultado, tmpOperar, "+")
			gen.AddExpression(t5, t5, "1", "-")
			gen.AddExpression(tpeso, tpeso, "10", "*")
			gen.AddGoto(LRec)

			gen.AddLabel(L9)
			tmpOperar2 := gen.NewTemp()
			gen.AddExpression(tmpComparar, tmpComparar, "48", "-")
			gen.AddExpression(tmpOperar2, tmpComparar, tpeso, "*")
			gen.AddExpression(tresultado, tresultado, tmpOperar2, "+")
			gen.AddExpression(t5, t5, "1", "-")
			gen.AddExpression(tpeso, tpeso, "10", "*")
			gen.AddGoto(LRec)

			gen.AddLabel(L5)
			gen.AddComment("No se puede convertir la cadena a INT")
			gen.AddPrintf("c", "78")
			gen.AddPrintf("c", "111")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "115")
			gen.AddPrintf("c", "101")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "112")
			gen.AddPrintf("c", "117")
			gen.AddPrintf("c", "101")
			gen.AddPrintf("c", "100")
			gen.AddPrintf("c", "101")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "99")
			gen.AddPrintf("c", "111")
			gen.AddPrintf("c", "110")
			gen.AddPrintf("c", "118")
			gen.AddPrintf("c", "101")
			gen.AddPrintf("c", "114")
			gen.AddPrintf("c", "116")
			gen.AddPrintf("c", "105")
			gen.AddPrintf("c", "114")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "97")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "73")
			gen.AddPrintf("c", "78")
			gen.AddPrintf("c", "84")
			gen.AddGoto(LSalir)
			gen.AddLabel(LSalir)

			val := environment.Value{Value: tresultado, IsTemp: true, Type: environment.INTEGER}
			return val
		}

	case "Float":
		{
			op1 = o.OpIzq.Execute(ast, env, gen)
			gen.AddComment("----FLOAT CAST---")
			t2 := gen.NewTemp()
			t3 := gen.NewTemp()
			t4 := gen.NewTemp()
			tpeso := gen.NewTemp()

			L1 := gen.NewLabel()
			L2 := gen.NewLabel()
			L3 := gen.NewLabel()
			L4 := gen.NewLabel()
			L5 := gen.NewLabel()

			gen.AddAssign(tpeso, "1")
			gen.AddAssign(t2, op1.Value)
			gen.AddAssign(t3, "0")
			gen.AddLabel(L4)
			gen.AddAssign(t4, "heap[(int)"+t2+"]")
			gen.AddIf(t4, "-1", "==", L1)
			gen.AddComment("viendo si solo numeros")
			gen.AddIf(t4, "48", "<", L5)
			gen.AddIf(t4, "57", ">", L5)

			gen.AddGoto(L2)
			gen.AddLabel(L1)

			gen.AddGoto(L3)
			gen.AddLabel(L2)
			gen.AddAssign(t3, t2)
			gen.AddExpression(t2, t2, "1", "+")
			gen.AddGoto(L4)
			gen.AddLabel(L3)

			t5 := gen.NewTemp()
			t6 := gen.NewTemp()
			L6 := gen.NewLabel()
			L7 := gen.NewLabel()
			tresultado := gen.NewTemp()
			LSalir := gen.NewLabel()

			gen.AddAssign(t5, t3)
			LRec := gen.NewLabel()

			gen.AddLabel(LRec)
			gen.AddAssign(t6, "heap[(int)"+t5+"]")
			gen.AddIf(t6, "-2", "==", L6)
			gen.AddIf(t6, "-3", "==", L5)
			gen.AddIf(t6, "-1", "==", LSalir)
			gen.AddIf(t5, "-1", "==", LSalir)
			gen.AddGoto(L7)

			gen.AddLabel(L6)
			gen.AddExpression(t5, t5, "1", "-")
			gen.AddAssign(tresultado, "heap[(int)"+t5+"]")
			gen.AddGoto(LSalir)

			gen.AddLabel(L7)
			tmpComparar := gen.NewTemp()
			L8 := gen.NewLabel()
			L9 := gen.NewLabel()
			gen.AddAssign(tmpComparar, "heap[(int)"+t5+"]")
			gen.AddIf(tpeso, "1", "==", L8)
			gen.AddGoto(L9)
			gen.AddLabel(L8)

			tmpOperar := gen.NewTemp()
			gen.AddExpression(tmpComparar, tmpComparar, "48", "-")
			gen.AddExpression(tmpOperar, tmpComparar, "1", "*")
			gen.AddExpression(tresultado, tresultado, tmpOperar, "+")
			gen.AddExpression(t5, t5, "1", "-")
			gen.AddExpression(tpeso, tpeso, "10", "*")
			gen.AddGoto(LRec)

			gen.AddLabel(L9)
			tmpOperar2 := gen.NewTemp()
			gen.AddExpression(tmpComparar, tmpComparar, "48", "-")
			gen.AddExpression(tmpOperar2, tmpComparar, tpeso, "*")
			gen.AddExpression(tresultado, tresultado, tmpOperar2, "+")
			gen.AddExpression(t5, t5, "1", "-")
			gen.AddExpression(tpeso, tpeso, "10", "*")
			gen.AddGoto(LRec)

			gen.AddLabel(L5)
			gen.AddComment("No se puede convertir la cadena a INT")
			gen.AddPrintf("c", "78")
			gen.AddPrintf("c", "111")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "115")
			gen.AddPrintf("c", "101")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "112")
			gen.AddPrintf("c", "117")
			gen.AddPrintf("c", "101")
			gen.AddPrintf("c", "100")
			gen.AddPrintf("c", "101")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "99")
			gen.AddPrintf("c", "111")
			gen.AddPrintf("c", "110")
			gen.AddPrintf("c", "118")
			gen.AddPrintf("c", "101")
			gen.AddPrintf("c", "114")
			gen.AddPrintf("c", "116")
			gen.AddPrintf("c", "105")
			gen.AddPrintf("c", "114")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "97")
			gen.AddPrintf("c", "32")
			gen.AddPrintf("c", "73")
			gen.AddPrintf("c", "78")
			gen.AddPrintf("c", "84")
			gen.AddGoto(LSalir)
			gen.AddLabel(LSalir)

			val := environment.Value{Value: tresultado, IsTemp: true, Type: environment.FLOAT}
			return val
		}

	case "String":
		// todo string cast
	}
	gen.AddBr()
	return environment.Value{}
}
