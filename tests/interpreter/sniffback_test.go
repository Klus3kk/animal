package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputSniffback(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_SniffbackDirectValue(t *testing.T) {
	code := `
5 meow 3 sniffback
`
	result, err := interpretInputSniffback(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 8 {
		t.Errorf("Expected result 8, got %v", result)
	}
}

func TestInterpreter_SniffbackInsideGrowl(t *testing.T) {
	code := `
growl true {
    10 sniffback
}
`
	result, err := interpretInputSniffback(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 10 {
		t.Errorf("Expected result 10 from growl sniffback, got %v", result)
	}
}

func TestInterpreter_SniffbackInsideFunction(t *testing.T) {
	code := `
howl double(x) {
    x meow x sniffback
}
double(6)
`
	result, err := interpretInputSniffback(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 12 {
		t.Errorf("Expected 12 from double(6), got %v", result)
	}
}
