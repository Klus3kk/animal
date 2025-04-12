package animal_test

import (
	"animal/animal/v2"
	"testing"
)

func TestArgs_PassedIntoScript(t *testing.T) {
	context := &animal.Context{
		DisplayName:  "<test>",
		Symbol_Table: animal.NewSymbolTable(),
	}
	context.Symbol_Table.set("__args__", []interface{}{"foo", "bar", "baz"})

	result, err := animal.Run(`__args__`, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}

	args, ok := result.([]interface{})
	if !ok {
		t.Fatalf("Expected list result, got: %T", result)
	}
	if len(args) != 3 || args[0] != "foo" || args[2] != "baz" {
		t.Errorf("Expected [foo, bar, baz], got: %v", result)
	}
}

func TestInterpreter_Arithmetic(t *testing.T) {
	result, err := animal.Run(`5 meow 3`, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if result != float64(8) {
		t.Errorf("Expected 8, got %v", result)
	}
}

func TestInterpreter_AssignmentAndAccess(t *testing.T) {
	code := `
x -> 10
x`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if result != float64(10) {
		t.Errorf("Expected 10, got %v", result)
	}
}

func TestInterpreter_GrowlSniffWag(t *testing.T) {
	code := `
x -> 4
growl x > 5 {
    99 sniffback
} sniff x == 4 {
    42 sniffback
} wag {
    0 sniffback
}`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if result != float64(42) {
		t.Errorf("Expected 42, got %v", result)
	}
}

func TestInterpreter_FunctionWithSniffback(t *testing.T) {
	code := `
howl sum(a, b) {
    a meow b sniffback
}
sum(7, 5)`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if result != float64(12) {
		t.Errorf("Expected 12, got %v", result)
	}
}

func TestInterpreter_LeapLoopSum(t *testing.T) {
	code := `
total -> 0
leap i from 0 to 3 {
    total -> total meow i
}
total`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if result != float64(3) {
		t.Errorf("Expected 3, got %v", result)
	}
}
