package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputGrowl(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_GrowlTrueCondition(t *testing.T) {
	result, err := interpretInputGrowl(t, `
growl true {
    42 sniffback
}
`)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 42 {
		t.Errorf("Expected 42, got %v", result)
	}
}

func TestInterpreter_GrowlSniffCondition(t *testing.T) {
	result, err := interpretInputGrowl(t, `
growl false {
    1 sniffback
}
sniff true {
    2 sniffback
}
`)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 2 {
		t.Errorf("Expected 2 from sniff, got %v", result)
	}
}

func TestInterpreter_GrowlWagCondition(t *testing.T) {
	result, err := interpretInputGrowl(t, `
growl false {
    1 sniffback
}
wag {
    3 sniffback
}
`)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 3 {
		t.Errorf("Expected 3 from wag, got %v", result)
	}
}

func TestInterpreter_GrowlNoMatchNoWag(t *testing.T) {
	result, err := interpretInputGrowl(t, `
growl false {
    1 sniffback
}
sniff false {
    2 sniffback
}
`)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil when no condition matched, got %v", result)
	}
}
