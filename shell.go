package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
