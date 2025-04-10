package animal_test

import (
	"animal/animal/v2"
	"testing"
)

func TestRoar_String(t *testing.T) {
	code := `
roar "Hello world"`
	_, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
}

func TestRoar_NumberExpr(t *testing.T) {
	code := `
x -> 5 meow 10
roar x`
	_, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
}

func TestRoar_MultipleValues(t *testing.T) {
	code := `
roar "Total:", 5 meow 3`
	_, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
}

func TestRoar_Empty(t *testing.T) {
	code := `
roar`
	_, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Expected no error on empty roar, got: %v", err)
	}
}
