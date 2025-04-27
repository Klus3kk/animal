package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputTryCatch(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_TryCatchErrorNameAvailable(t *testing.T) {
	code := `
*[ 
    *{ "something bad" }*
]*
*(
    _error sniffback
)*
`
	result, err := interpretInputTryCatch(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if str, ok := result.(string); !ok || str != "something bad" {
		t.Errorf("Expected thrown error 'something bad', got %v", result)
	}
}
