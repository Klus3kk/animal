package main

import (
	"animal/core"
	"animal/core/std"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const AnimalVersion = "0.1.1"

func debugParse(input string) {
	lexer := core.NewLexer("<stdin>", input)
	tokens, err := lexer.MakeTokens()
	if err != nil {
		fmt.Println("Lexer error:", err)
		return
	}
	fmt.Println("Tokens:", tokens)

	parser := core.NewParser(tokens)
	parseResult := parser.Parse()
	if parseResult.Error != "" {
		fmt.Println("Parser error:", parseResult.Error)
	} else {
		fmt.Println("AST:", parseResult.Node)
	}
}

func main() {
	args := os.Args[1:]

	// Handle --help and --version flags
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			fmt.Println("Animal Programming Language")
			fmt.Println()
			fmt.Println("Usage:")
			fmt.Println("  animal [file.anml] [args...]    Run a file")
			fmt.Println("  animal --repl                   Start interactive REPL mode")
			fmt.Println("  animal --help                   Show this help message")
			fmt.Println("  animal --version                Show version information")
			fmt.Println("  animal --debug                  Enable debug mode")
			fmt.Println("  animal --time                   Measure execution time")
			os.Exit(0)
		}
		if arg == "--version" {
			fmt.Println("Animal Programming Language Version:", AnimalVersion)
			os.Exit(0)
		}
	}

	// Handle --debug and --time flags
	for idx, arg := range args {
		if arg == "--debug" {
			core.Debug = true
			args = append(args[:idx], args[idx+1:]...)
			break
		}
		if arg == "--time" {
			core.TimeEnabled = true
			args = append(args[:idx], args[idx+1:]...)
			break
		}
	}

	// Create global context
	globalSymbolTable := core.NewSymbolTable()
	std.RegisterStandardLibrary(globalSymbolTable)
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	// Inject CLI arguments into __args__
	if len(args) > 0 {
		interfaceArgs := make([]interface{}, len(args))
		for i, v := range args {
			interfaceArgs[i] = v
		}
		context.Symbol_Table.Set("__args__", interfaceArgs)
	}

	// If first argument is a .anml file, run it
	if len(args) >= 1 && strings.HasSuffix(args[0], ".anml") {
		codeBytes, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Printf("\033[31mError reading file: %v\033[0m\n", err)
			return
		}
		code := string(codeBytes)
		code = strings.TrimPrefix(code, "\uFEFF") // remove BOM if present

		// Skip shebang if present
		if strings.HasPrefix(code, "#!") {
			lines := strings.SplitN(code, "\n", 2)
			if len(lines) > 1 {
				code = lines[1]
			} else {
				code = ""
			}
		}

		start := time.Now()

		_, err = core.CustomRun(code, args[0], context)

		if err != nil {
			fmt.Printf("\033[31m%v\033[0m\n", err)
		}

		if core.TimeEnabled {
			elapsed := time.Since(start)
			fmt.Printf("Total execution time: %v\n", elapsed)
		}

		return
	}

	// Default: REPL Mode
	fmt.Printf("\033[36mAnimal %s\033[0m\n", AnimalVersion)
	fmt.Println("Type 'exit' to quit. Type '%debug' to toggle debug mode. Type '%time' to toggle timing mode.")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		inputLower := strings.ToLower(input)

		if inputLower == "exit" {
			fmt.Println("Exiting...")
			return
		}

		if inputLower == "%debug" {
			core.Debug = !core.Debug
			fmt.Println("Debug mode:", core.Debug)
			continue
		}

		if inputLower == "%time" {
			core.TimeEnabled = !core.TimeEnabled
			fmt.Println("Timing mode:", core.TimeEnabled)
			continue
		}

		var start time.Time
		if core.TimeEnabled {
			start = time.Now()
		}

		result, err := core.CustomRun(input, "<stdin>", context)

		if core.TimeEnabled {
			elapsed := time.Since(start)
			fmt.Printf("Execution time: %v\n", elapsed)
		}

		if err != nil {
			fmt.Printf("\033[31m%v\033[0m\n", err)
		} else if result != nil {
			fmt.Println(result)
		}
	}
}
