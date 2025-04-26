package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputLeap(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_LeapSimple(t *testing.T) {
	code := `
sum -> 0
leap i from 0 to 5 {
    sum -> sum meow i
}
roar sum
`
	result, err := interpretInputLeap(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := 0 + 1 + 2 + 3 + 4 // leap is exclusive: to 5 (not including 5)
	if num, ok := result.(float64); !ok || int(num) != expected {
		t.Errorf("Expected sum %d, got %v", expected, result)
	}
}

func TestInterpreter_LeapBreakWithWhimper(t *testing.T) {
	code := `
sum -> 0
leap i from 0 to 5 {
    growl i == 3 {
        whimper
    }
    sum -> sum meow 1
}
roar sum
`
	result, err := interpretInputLeap(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || int(num) != 3 {
		t.Errorf("Expected sum 3 after whimper break, got %v", result)
	}
}

func TestInterpreter_LeapContinueWithHiss(t *testing.T) {
	code := `
sum -> 0
leap i from 0 to 5 {
    growl i == 2 {
        hiss
    }
    sum -> sum meow 1
}
roar sum
`
	result, err := interpretInputLeap(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if num, ok := result.(float64); !ok || int(num) != 4 {
		t.Errorf("Expected sum 4 (one hissed iteration skipped), got %v", result)
	}
}
