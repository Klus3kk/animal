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
throw *{ "something went wrong" }*
`
	_, err := interpretInputThrow(t, code)
	if err == nil {
		t.Fatalf("Expected an error to be thrown, got none")
	}

	if err.Error() != "something went wrong" {
		t.Errorf("Expected 'something went wrong', got %v", err)
	}
}

func TestInterpreter_ThrowInsideTryCatch(t *testing.T) {
	code := `
*[ 
    throw *{ "bad thing" }*
]* *(
    "handled" sniffback
)*
`
	result, err := interpretInputThrow(t, code)
	if err != nil {
		t.Fatalf("Unexpected error inside catch: %v", err)
	}

	if str, ok := result.(string); !ok || str != "handled" {
		t.Errorf("Expected 'handled' from catch block, got %v", result)
	}
}
