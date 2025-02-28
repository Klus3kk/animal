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

func main() {
	// debugParse("x -> 0")
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

		// Now run returns an AST instead of tokens
		ast, err := run(input, "<stdin>")
		if err != nil {
			fmt.Println(err)
		} else if ast != nil {
			fmt.Println(ast) // Print the AST
		}
	}
}
