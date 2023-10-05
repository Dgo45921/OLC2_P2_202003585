package environment

import (
	"html"
	"strconv"
)

type CustomSemanticError struct {
	line, column int
	msg          string
	ttype        string
}

type AST struct {
	Instructions []interface{}
	Print        string
	Errors       []CustomSemanticError
	Symbols      map[string]Symbol
	FuncSymbol   map[string]FunctionSymbol
}

func (a *AST) GetPrint() string {
	return a.Print
}

func (a *AST) SetPrint(ToPrint string) {
	a.Print = a.Print + ToPrint
}

func (a *AST) GetErrors() string {
	response := "<tr><td colspan=\"4\" bgcolor=\"lightgrey\"><b>Semanticos</b></td></tr>"
	for _, err := range a.Errors {
		response += "<tr><td>Semantico</td><td>" + html.EscapeString(err.msg) + "</td><td>" + strconv.Itoa(err.line) + "</td><td>" + strconv.Itoa(err.column) + "</td></tr>"
	}

	return response
}

func (a *AST) SetError(line int, col int, des string) {
	err := CustomSemanticError{
		line:   line,
		column: col,
		msg:    des,
		ttype:  "Semantico",
	}
	a.SetPrint("Error: " + des + "\n")
	a.Errors = append(a.Errors, err)
}

func (a *AST) SaveSymbol(id string, symbol Symbol) {
	a.Symbols[id] = symbol
}

func (a *AST) SaveFunction(id string, symbol FunctionSymbol) {
	a.FuncSymbol[id] = symbol
}

func (a *AST) GetSymbolTable() string {
	vizcode := "digraph G {\n  node [shape=plaintext];\n  labelloc=\"t\";\n  label=<\n    <table border=\"1\" cellspacing=\"0\" cellpadding=\"10\" >\n      <tr>\n        <td bgcolor=\"lightgrey\" align=\"center\"><b>ID</b></td>\n        <td bgcolor=\"lightgrey\" align=\"center\"><b>Tipo símbolo</b></td>\n        <td bgcolor=\"lightgrey\" align=\"center\"><b>Tipo dato</b></td>\n        <td bgcolor=\"lightgrey\" align=\"center\"><b>Ámbito</b></td>\n        <td bgcolor=\"lightgrey\" align=\"center\"><b>Línea</b></td>\n        <td bgcolor=\"lightgrey\" align=\"center\"><b>Columna</b></td>\n      </tr>"
	for key, element := range a.Symbols {
		if element.Const {
			vizcode += "<tr>\n<td>" + key + "</td>\n   <td>Constante</td>\n    <td>" + element.GetType() + "</td>\n   <td>" + element.GetScopeType() + "</td>\n  <td>" + strconv.Itoa(element.Lin) + "</td>\n        <td>" + strconv.Itoa(element.Col) + "</td>\n      </tr>"
		} else {
			if element.Type == STRUCT_DEF {
				vizcode += "<tr>\n<td>" + key + "</td>\n   <td>Definicion struct</td>\n    <td>" + element.GetType() + "</td>\n   <td>" + element.GetScopeType() + "</td>\n  <td>" + strconv.Itoa(element.Lin) + "</td>\n        <td>" + strconv.Itoa(element.Col) + "</td>\n      </tr>"

			} else {
				vizcode += "<tr>\n<td>" + key + "</td>\n   <td>Variable</td>\n    <td>" + element.GetType() + "</td>\n   <td>" + element.GetScopeType() + "</td>\n  <td>" + strconv.Itoa(element.Lin) + "</td>\n        <td>" + strconv.Itoa(element.Col) + "</td>\n      </tr>"

			}
		}
	}

	for key, element := range a.FuncSymbol {
		vizcode += "<tr>\n<td>" + key + "</td>\n   <td>Funcion</td>\n    <td>" + element.GetType() + "</td>\n   <td>" + "GLOBAL" + "</td>\n  <td>" + strconv.Itoa(element.Lin) + "</td>\n        <td>" + strconv.Itoa(element.Col) + "</td>\n      </tr>"

	}

	vizcode += "    </table>\n  >;\n}\n"
	return vizcode
}
