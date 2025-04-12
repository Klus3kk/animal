package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func debugParse(input string) {
	lexer := NewLexer("<stdin>", input)
	tokens, err := lexer.make_tokens()
	if err != nil {
		fmt.Println("Lexer error:", err)
		return
	}
	fmt.Println("Tokens:", tokens)

	parser := NewParser(tokens)
	parseResult := parser.parse()
	if parseResult.Error != "" {
		fmt.Println("Parser error:", parseResult.Error)
	} else {
		fmt.Println("AST:", parseResult.Node)
	}
}

func customRun(text string, fn string, context *Context) (interface{}, error) {
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

	// Create an interpreter and use the provided context
	interpreter := Interpreter{}
	result := interpreter.visit(parseResult.Node, context)

	if result.Error != nil {
		return nil, result.Error
	}

	return result.Value, nil
}

func main() {
	args := os.Args[1:]

	// Create a global symbol table to be used across inputs
	globalSymbolTable := NewSymbolTable()
	context := &Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	// Inject CLI arguments as __args__ (converted to []interface{})
	if len(args) > 0 {
		interfaceArgs := make([]interface{}, len(args))
		for i, v := range args {
			interfaceArgs[i] = v
		}
		context.Symbol_Table.set("__args__", interfaceArgs)
	}

	// Run file mode: ./animal script.anml arg1 arg2 ...
	if len(args) >= 1 && strings.HasSuffix(args[0], ".anml") {
		codeBytes, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		code := string(codeBytes)
		code = strings.TrimPrefix(code, "\uFEFF") // remove BOM if present

		// Skip shebang line if present
		if strings.HasPrefix(code, "#!") {
			lines := strings.SplitN(code, "\n", 2)
			if len(lines) > 1 {
				code = lines[1]
			} else {
				code = ""
			}
		}

		_, err = customRun(code, args[0], context)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Default: REPL mode
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		inputLower := strings.ToLower(input)

		if inputLower == "exit" {
			fmt.Println("Exiting...")
			return
		}

		result, err := customRun(input, "<stdin>", context)
		if err != nil {
			fmt.Println(err)
		} else if result != nil {
			fmt.Println(result)
		}
	}
}
