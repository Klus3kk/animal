package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputFunction(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_FunctionSimpleCall(t *testing.T) {
	code := `
howl greet(name) {
    roar "Hello", name
}
greet("Animal")
`
	result, err := interpretInputFunction(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil after function call (roar inside), got %v", result)
	}
}

func TestInterpreter_FunctionReturnValue(t *testing.T) {
	code := `
howl add(a, b) {
    a meow b sniffback
}
add(3, 4)
`
	result, err := interpretInputFunction(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 7 {
		t.Errorf("Expected result 7, got %v", result)
	}
}

func TestInterpreter_FunctionMultipleCalls(t *testing.T) {
	code := `
howl square(x) {
    x moo x sniffback
}
square(2)
square(5)
`
	result, err := interpretInputFunction(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should return the last call result (square(5) = 25)
	if num, ok := result.(float64); !ok || num != 25 {
		t.Errorf("Expected result 25, got %v", result)
	}
}

func TestInterpreter_FunctionArgumentMismatch(t *testing.T) {
	code := `
howl greet(name) {
    roar "Hello", name
}
greet()
`
	_, err := interpretInputFunction(t, code)
	if err == nil {
		t.Fatalf("Expected error due to wrong number of arguments")
	}
}
