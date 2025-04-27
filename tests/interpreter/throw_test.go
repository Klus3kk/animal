package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputThrow(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_ThrowSimpleError(t *testing.T) {
	code := `
*{ "something went wrong" }*
`
	_, err := interpretInputThrow(t, code)
	if err == nil {
		t.Fatalf("Expected an error to be thrown, got none")
	}

	if err.Error() != "something went wrong" {
		t.Errorf("Expected 'something went wrong', got %v", err)
	}
}
