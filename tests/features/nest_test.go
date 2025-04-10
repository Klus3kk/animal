package animal_test

import (
	"animal/animal/v2"
	"reflect"
	"testing"
)

func TestNest_BasicFields(t *testing.T) {
	code := `
nest Dog {
    name
}

d -> Dog()
d.name -> "Burek"
d.name`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	if result != "Burek" {
		t.Errorf("Expected 'Burek', got %v", result)
	}
}

func TestNest_MethodCall(t *testing.T) {
	code := `
nest Dog {
    name
    howl greet() {
        "woof " purr this.name sniffback
    }
}

d -> Dog()
d.name -> "Max"
d.greet()`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := "woof Max"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNest_FieldSetAndGet(t *testing.T) {
	code := `
nest Cat {
    name
    age
}

c -> Cat()
c.name -> "Mruczek"
c.age -> 7
[c.name, c.age]`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := []interface{}{"Mruczek", float64(7)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestNest_MethodAccessesThis(t *testing.T) {
	code := `
nest Bird {
    name
    howl sing() {
        "tweet " purr this.name sniffback
    }
}

b -> Bird()
b.name -> "Chirpy"
b.sing()`
	result, err := animal.Run(code, "<test>")
	if err != nil {
		t.Fatalf("Runtime error: %v", err)
	}
	expected := "tweet Chirpy"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
