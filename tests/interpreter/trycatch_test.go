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

func TestInterpreter_TryCatchNoError(t *testing.T) {
	code := `
*[ 
    5 meow 3 sniffback
]* *(
    "caught" sniffback
)*
`
	result, err := interpretInputTryCatch(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 8 {
		t.Errorf("Expected result 8 (normal body execution), got %v", result)
	}
}

func TestInterpreter_TryCatchWithError(t *testing.T) {
	code := `
*[ 
    5 drone 0 sniffback
]* *(
    "caught error" sniffback
)*
`
	result, err := interpretInputTryCatch(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if str, ok := result.(string); !ok || str != "caught error" {
		t.Errorf("Expected 'caught error', got %v", result)
	}
}

func TestInterpreter_TryCatchErrorNameAvailable(t *testing.T) {
	code := `
*[ 
    throw *{ "something bad" }*
]* *(
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
