package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputNest(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_NestCreateAndAccessField(t *testing.T) {
	code := `
nest Dog {
    name
}
d -> Dog()
d.name -> "Buddy"
d.name
`
	result, err := interpretInputNest(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if str, ok := result.(string); !ok || str != "Buddy" {
		t.Errorf("Expected 'Buddy' from d.name, got %v", result)
	}
}

func TestInterpreter_NestMethodCall(t *testing.T) {
	code := `
nest Dog {
    name
    howl bark() {
        "woof" sniffback
    }
}
d -> Dog()
d.bark()
`
	result, err := interpretInputNest(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if str, ok := result.(string); !ok || str != "woof" {
		t.Errorf("Expected 'woof' from d.bark(), got %v", result)
	}
}

func TestInterpreter_NestAccessUnknownField(t *testing.T) {
	code := `
nest Dog {
    name
}
d -> Dog()
d.age
`
	_, err := interpretInputNest(t, code)
	if err == nil {
		t.Fatalf("Expected error for accessing unknown field 'age'")
	}
}
