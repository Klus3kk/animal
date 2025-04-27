package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputLeap(t *testing.T, input string) (*core.SymbolTable, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	_, err := core.CustomRun(input, "<stdin>", context)
	return globalSymbolTable, err
}

func TestInterpreter_LeapSimple(t *testing.T) {
	code := `
sum -> 0
leap i from 0 to 5 {
    sum -> sum meow i
}
`
	table, err := interpretInputLeap(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	sum, _ := table.Get("sum")
	if int(sum.Value.(float64)) != 10 {
		t.Errorf("Expected sum 10, got %v", sum.Value)
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
`
	table, err := interpretInputLeap(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	sum, _ := table.Get("sum")
	if int(sum.Value.(float64)) != 3 {
		t.Errorf("Expected sum 3 after whimper break, got %v", sum.Value)
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
`
	table, err := interpretInputLeap(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	sum, _ := table.Get("sum")
	if int(sum.Value.(float64)) != 4 {
		t.Errorf("Expected sum 4 (one hissed iteration skipped), got %v", sum.Value)
	}
}
