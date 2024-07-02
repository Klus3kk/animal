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

		tokens, err := run(input, "<stdin>")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Tokens:")
			for _, token := range tokens {
				fmt.Println(token)
			}
		}
	}
}
