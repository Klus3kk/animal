package core

import (
	"fmt"
	"strings"
)


var (
	OutputBuffer  strings.Builder
	CaptureOutput = false
)

func Print(args ...interface{}) {
	text := fmt.Sprint(args...)
	if CaptureOutput {
		OutputBuffer.WriteString(text + "\n")
	} else {
		fmt.Println(args...)
	}
}

func GetCapturedOutput() string {
	return OutputBuffer.String()
}

func ClearCapturedOutput() {
	OutputBuffer.Reset()
}

func CustomRun(text string, fn string, context *Context) (interface{}, error) {
	lexer := NewLexer(fn, text)
	tokens, err := lexer.make_tokens()
	if err != nil {
		return nil, err
	}

	parser := NewParser(tokens)
	parseResult := parser.parse()

	if parseResult.Error != "" {
		return nil, fmt.Errorf(parseResult.Error)
	}

	interpreter := Interpreter{}
	result := interpreter.visit(parseResult.Node, context)

	if result.Error != nil {
		return nil, result.Error
	}

	return result.Value, nil
}

func Run(text string, fn string) (interface{}, error) {
	return run(text, fn)
}

// RUN //
func run(text string, fn string) (interface{}, error) {
	// Create a new global symbol table
	globalSymbolTable := NewSymbolTable()

	// Make the context
	context := &Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	// Initialize the lexer and generate tokens
	lexer := NewLexer(fn, text)
	tokens, err := lexer.make_tokens()
	if err != nil {
		return nil, err
	}

	// Parse the tokens to generate the AST
	parser := NewParser(tokens)
	parseResult := parser.parse()

	if parseResult.Error != "" {
		return nil, fmt.Errorf(parseResult.Error)
	}

	// Create an interpreter with the SAME context
	interpreter := Interpreter{}
	result := interpreter.visit(parseResult.Node, context)

	if result.Error != nil {
		return nil, result.Error
	}

	return result.Value, nil
}
