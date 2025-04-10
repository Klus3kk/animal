package animal_test

import (
	"animal/animal/v2"
	"testing"
)

func TestFunction_SimpleAddition(t *testing.T) {
	code := `
howl add(a, b) {
    a meow b sniffback
}
add(5, 7)`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := float64(12)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFunction_NestedCalls(t *testing.T) {
	code := `
howl square(x) {
    x moo x sniffback
}
howl sum_of_squares(a, b) {
    square(a) meow square(b) sniffback
}
sum_of_squares(3, 4)`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := float64(25)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFunction_ReturnEarly(t *testing.T) {
	code := `
howl test() {
    10 sniffback
    99
}
test()`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := float64(10)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFunction_WithListArg(t *testing.T) {
	code := `
howl getFirst(lst) {
    lst sniffback
}
getFirst([1, 2, 3])`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if result == nil || len(result.([]interface{})) != 3 {
		t.Errorf("Expected full list, got %v", result)
	}
}
