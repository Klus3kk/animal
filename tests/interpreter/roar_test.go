package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInput(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_RoarPrintSingleValue(t *testing.T) {
	result, err := interpretInput(t, `roar 5`)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil after roar, got %v", result)
	}
}

func TestInterpreter_RoarPrintMultipleValues(t *testing.T) {
	result, err := interpretInput(t, `roar "hello", "world", 42`)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil after multiple roar, got %v", result)
	}
}

func TestInterpreter_RoarEmpty(t *testing.T) {
	result, err := interpretInput(t, `roar`)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil after empty roar, got %v", result)
	}
}

func TestInterpreter_RoarExpression(t *testing.T) {
	result, err := interpretInput(t, `roar 3 meow 4`)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil after roaring an expression, got %v", result)
	}
}
