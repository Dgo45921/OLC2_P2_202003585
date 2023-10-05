package controllers

import (
	"fmt"
	"html"
	"strconv"
)

func getVizCode(lexerErrors *CustomLexicalErrorListener, parserErrors *CustomSyntaxErrorListener, semanticstring string) string {
	vizcode := "digraph G {\n  node [shape=plaintext];\n  labelloc=\"t\";\n  label=\"Reporte de Errores\";\n  \n  a [label=<\n    <table border=\"0\" cellborder=\"1\" cellspacing=\"0\">\n      <tr><td bgcolor=\"lightgrey\"><b>Tipo</b></td><td bgcolor=\"lightgrey\"><b>Descripción</b></td><td bgcolor=\"lightgrey\"><b>Línea</b></td><td bgcolor=\"lightgrey\"><b>Columna</b></td></tr>"
	vizcode += "<tr><td colspan=\"4\" bgcolor=\"lightgrey\"><b>Lexicos</b></td></tr>"
	for _, err := range lexerErrors.Errors {
		vizcode += "<tr><td>Lexico</td><td>" + html.EscapeString(err.msg) + "</td><td>" + strconv.Itoa(err.line) + "</td><td>"+ strconv.Itoa(err.column)  +"</td></tr>"
	}

	vizcode += "<tr><td colspan=\"4\" bgcolor=\"lightgrey\"><b>Sintacticos</b></td></tr>"
	for _, err := range parserErrors.Errors {
		vizcode += "<tr><td>Sintactico</td><td>" + html.EscapeString(err.msg) + "</td><td>" + strconv.Itoa(err.line) + "</td><td>"+ strconv.Itoa(err.column)  +"</td></tr>"
	}

	fmt.Println(lexerErrors.Errors)
	fmt.Println(parserErrors.Errors)

	vizcode += semanticstring
	vizcode += "</table>\n  >];\n}\n"

	return vizcode
}
