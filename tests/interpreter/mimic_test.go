package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputMimic(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_MimicMatchingCase(t *testing.T) {
	code := `
mimic 2 {
    1 -> "one" sniffback
    2 -> "two" sniffback
    _ -> "other" sniffback
}
`
	result, err := interpretInputMimic(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if str, ok := result.(string); !ok || str != "two" {
		t.Errorf("Expected 'two' match, got %v", result)
	}
}

func TestInterpreter_MimicDefaultCase(t *testing.T) {
	code := `
mimic 5 {
    1 -> "one" sniffback
    2 -> "two" sniffback
    _ -> "other" sniffback
}
`
	result, err := interpretInputMimic(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if str, ok := result.(string); !ok || str != "other" {
		t.Errorf("Expected 'other' from default case, got %v", result)
	}
}

func TestInterpreter_MimicNoCasesMatchNoDefault(t *testing.T) {
	code := `
mimic 7 {
    1 -> "one" sniffback
    2 -> "two" sniffback
}
`
	result, err := interpretInputMimic(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil when no cases match and no default, got %v", result)
	}
}
