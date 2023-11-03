package controllers

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"PY1/models"
	"PY1/parser"
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var lastGivencode = ""
var lastReturnedCode []string
var lexerErrors = &CustomLexicalErrorListener{}
var parserErrors = &CustomSyntaxErrorListener{}

// Ast create ast
var Ast environment.AST

type TreeShapeListener struct {
	*parser.BaseSwiftGrammarListener
	Code []interface{}
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello skrrr!")
	if err != nil {
		return
	}
}

// ========== optmize functions =============================
func validateLineRule1(line string) bool {
	patron := `^t\d+\s*=\s*(t\d+|\d+)\s*[+\-\*/]\s*(t\d+|\d+)$`
	regex := regexp.MustCompile(patron)
	flag := regex.MatchString(line)
	return flag
}

func getTokensRule1(line string) (string, string, string, string) {
	tempArr := strings.Split(line, " ")
	return tempArr[0], tempArr[2], tempArr[3], tempArr[4]
}

func cleanLineRule1(line string) string {
	noJump := strings.ReplaceAll(line, "\n", "")
	noTab := strings.ReplaceAll(noJump, "\t", "")
	newLine := strings.ReplaceAll(noTab, ";", "")
	return newLine
}

func Rule1(arr []string) []string { //Eliminación de instrucciones red. de carga y almacenamiento
	//se recorre el arreglo
	for i := 0; i < len(arr); i++ {
		//leyendo entrada
		line := cleanLineRule1(arr[i])
		//comprobando
		if validateLineRule1(line) {
			//obteniendo tokens
			target1, left1, op1, right1 := getTokensRule1(line)
			//continuando recorrido
			for j := i + 1; j < len(arr); j++ {
				//leyendo nueva entrada
				line2 := cleanLineRule1(arr[j])
				if validateLineRule1(line2) {
					target2, left2, op2, right2 := getTokensRule1(line2)
					//validación 1
					if target2 == target1 || target2 == left1 || target2 == right1 {
						break
					}
					//validación 2
					if left1+op1+right1 == left2+op2+right2 || left1+op1+right1 == right2+op2+left2 {
						//sustituir
						arr[j] = "\t" + target2 + " = " + target1 + ";\n"
						continue
					}
				}

			}
		}
	}
	return arr
}

// ==========================================================

func Parse(w http.ResponseWriter, r *http.Request) {
	lexerErrors = &CustomLexicalErrorListener{}
	parserErrors = &CustomSyntaxErrorListener{}
	Ast = environment.AST{}
	Ast.Symbols = make(map[string]environment.Value)
	Ast.FuncSymbol = make(map[string]environment.FunctionSymbol)
	// newCode is responsible to save the given input
	var newCode models.SourceCode
	// consoleResponse is responsible of returning all of the console logs
	var consoleResponse models.ConsoleResponse
	// getting the body from the request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ERROR")
	}

	// parsing the json
	json.Unmarshal(reqBody, &newCode)
	// printing the input
	fmt.Println(newCode.Code)
	lastGivencode = newCode.Code
	lastReturnedCode = []string{}

	err = writeSourceCodeFile("source.txt", lastGivencode)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//Entrada
	var code string = newCode.Code
	//Leyendo entrada
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	//creacion de parser
	p := parser.NewSwiftGrammarParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)
	p.BuildParseTrees = true
	tree := p.S()
	//listener
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code

	// creating env
	var newEnv = environment.NewEnvironment(nil, environment.GLOBAL)
	//create generator
	var Generator generator.Generator
	Generator = generator.NewGenerator()

	Generator.MainCode = true
	//ejecución
	for _, inst := range Code {
		inst.(interfaces.Instruction).Execute(&Ast, newEnv, &Generator)
	}
	Generator.GenerateFinalCode()
	var ConsoleOut = ""
	if len(Ast.Errors) == 0 {
		for _, item := range Generator.GetFinalCode() {
			ConsoleOut += item.(string)
			lastReturnedCode = append(lastReturnedCode, item.(string))
		}
	} else {
		ConsoleOut = "Hubieron errores, por favor revise el reporte de errores!\n"
	}
	fmt.Println(ConsoleOut)

	consoleResponse.Console = ConsoleOut
	json.NewEncoder(w).Encode(consoleResponse)
}

func GetOptimized(w http.ResponseWriter, r *http.Request) {

	var consoleResponse models.ConsoleResponse

	codeR1 := Rule1(lastReturnedCode)

	//salida
	var ConsoleOut = ""
	for _, item := range codeR1 {
		ConsoleOut += item
	}

	consoleResponse.Console = ConsoleOut
	json.NewEncoder(w).Encode(consoleResponse)
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}

func writeSourceCodeFile(filename, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetCST(w http.ResponseWriter, r *http.Request) {

	exePath, errr := os.Executable()
	if errr != nil {
		log.Fatal(errr)
	}

	currentDir := filepath.Dir(exePath)
	cmd := exec.Command("antlr4-parse", currentDir+"/grammarforcst/Grammar.g4", "s", "-gui", "source.txt")
	cmd.Dir = currentDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func GetErrors(w http.ResponseWriter, r *http.Request) {

	var consoleResponse models.DotResponse

	consoleResponse.DotCode = getVizCode(lexerErrors, parserErrors, Ast.GetErrors())
	json.NewEncoder(w).Encode(consoleResponse)

}

func GetSymbolTable(w http.ResponseWriter, r *http.Request) {

	var consoleResponse models.DotResponse

	consoleResponse.DotCode = Ast.GetSymbolTable()
	json.NewEncoder(w).Encode(consoleResponse)

}

type CustomSyntaxError struct {
	line, column int
	msg          string
	ttype        string
}

type CustomLexicalErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []CustomSyntaxError
}

func (c *CustomLexicalErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	newError := CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
		ttype:  "Lexico",
	}
	c.Errors = append(c.Errors, newError)
}

type CustomSyntaxErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []CustomSyntaxError
}

func (c *CustomSyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	newError := CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
		ttype:  "Sintactico",
	}
	c.Errors = append(c.Errors, newError)
}
