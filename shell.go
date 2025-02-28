package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
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
    reader := bufio.NewReader(os.Stdin)
    
    // Create a global symbol table to be used across all inputs
    globalSymbolTable := NewSymbolTable()
    context := &Context{
        DisplayName:  "<program>",
        Symbol_Table: globalSymbolTable,
    }
    
    for {
        fmt.Print(">> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        inputLower := strings.ToLower(input)

        if inputLower == "exit" {
            fmt.Println("Exiting...")
            return
        }

        // Customize the run function to use the persistent context
        result, err := customRun(input, "<stdin>", context)
        if err != nil {
            fmt.Println(err)
        } else if result != nil {
            fmt.Println(result)
        }
    }
}
