package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputPounce(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_PounceBasicLoop(t *testing.T) {
	code := `
x -> 3
pounce x > 0 {
    x -> x woof 1
}
roar x
`
	result, err := interpretInputPounce(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 0 {
		t.Errorf("Expected x to be 0 after pounce, got %v", result)
	}
}

func TestInterpreter_PounceBreakWithWhimper(t *testing.T) {
	code := `
x -> 5
pounce x > 0 {
    growl x == 3 {
        whimper
    }
    x -> x woof 1
}
roar x
`
	result, err := interpretInputPounce(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 3 {
		t.Errorf("Expected x to be 3 after whimper break, got %v", result)
	}
}

func TestInterpreter_PounceContinueWithHiss(t *testing.T) {
	code := `
x -> 5
y -> 0
pounce x > 0 {
    x -> x woof 1
    growl x == 3 {
        hiss
    }
    y -> y meow 1
}
roar y
`
	result, err := interpretInputPounce(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 4 {
		t.Errorf("Expected y to be 4 (skipped increment once), got %v", result)
	}
}
