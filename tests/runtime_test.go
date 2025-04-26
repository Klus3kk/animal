package tests

import (
	"animal/core"
	"testing"
)

func interpretInputRuntime(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestRuntime_DivisionByZero(t *testing.T) {
	code := `
10 drone 0
`
	_, err := interpretInputRuntime(t, code)
	if err == nil {
		t.Fatalf("Expected division by zero error, got none")
	}
}

func TestRuntime_UndefinedVariable(t *testing.T) {
	code := `
roar z
`
	_, err := interpretInputRuntime(t, code)
	if err == nil {
		t.Fatalf("Expected undefined variable error, got none")
	}
}

func TestRuntime_TypeMismatchAssignment(t *testing.T) {
	code := `
x: int -> 5
x -> "string"
`
	_, err := interpretInputRuntime(t, code)
	if err == nil {
		t.Fatalf("Expected type mismatch error, got none")
	}
}

func TestRuntime_InvalidListIndex(t *testing.T) {
	code := `
l -> [1,2,3]
roar l[5]
`
	_, err := interpretInputRuntime(t, code)
	if err == nil {
		t.Fatalf("Expected list index out of bounds error, got none")
	}
}

func TestRuntime_InvalidFunctionCall(t *testing.T) {
	code := `
5(3)
`
	_, err := interpretInputRuntime(t, code)
	if err == nil {
		t.Fatalf("Expected function call error on non-function, got none")
	}
}
