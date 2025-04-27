package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputPounce(t *testing.T, input string) (*core.SymbolTable, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	_, err := core.CustomRun(input, "<stdin>", context)
	return globalSymbolTable, err
}

func TestInterpreter_PounceBasicLoop(t *testing.T) {
	code := `
x -> 3
pounce x > 0 {
    x -> x woof 1
}
`
	table, err := interpretInputPounce(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	x, _ := table.Get("x")
	if int(x.Value.(float64)) != 0 {
		t.Errorf("Expected x to be 0 after pounce, got %v", x.Value)
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
`
	table, err := interpretInputPounce(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	x, _ := table.Get("x")
	if int(x.Value.(float64)) != 3 {
		t.Errorf("Expected x to be 3 after whimper break, got %v", x.Value)
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
`
	table, err := interpretInputPounce(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	y, _ := table.Get("y")
	if int(y.Value.(float64)) != 4 {
		t.Errorf("Expected y to be 4 (skipped increment once), got %v", y.Value)
	}
}
