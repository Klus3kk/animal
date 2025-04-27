package interpreter

import (
	"animal/core"
	"testing"
)

func interpretInputListMethods(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_ListSniffAppend(t *testing.T) {
	code := `
l -> [1, 2]
l.sniff(3)
l
`
	result, err := interpretInputListMethods(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	list, ok := result.([]interface{})
	if !ok || len(list) != 3 || list[2].(float64) != 3 {
		t.Errorf("Expected list [1 2 3], got %v", result)
	}
}

//func TestInterpreter_ListHowlFindIndex(t *testing.T) {
//	code := `
//l -> [10, 20, 30]
//l.howl(20)
//`
//	result, err := interpretInputListMethods(t, code)
//	if err != nil {
//		t.Fatalf("Unexpected error: %v", err)
//	}
//
//	if idx, ok := result.(float64); !ok || int(idx) != 1 {
//		t.Errorf("Expected index 1 for 20, got %v", result)
//	}
//}

func TestInterpreter_ListWagLength(t *testing.T) {
	code := `
l -> [5, 6, 7, 8]
l.wag()
`
	result, err := interpretInputListMethods(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if length, ok := result.(float64); !ok || int(length) != 4 {
		t.Errorf("Expected length 4, got %v", result)
	}
}

func TestInterpreter_ListSnarlReverse(t *testing.T) {
	code := `
l -> [1, 2, 3]
l.snarl()
l
`
	result, err := interpretInputListMethods(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	list, ok := result.([]interface{})
	if !ok || list[0].(float64) != 3 {
		t.Errorf("Expected reversed list [3 2 1], got %v", result)
	}
}

func TestInterpreter_ListProwlShuffle(t *testing.T) {
	code := `
l -> [1, 2, 3, 4, 5]
l.prowl()
l
`
	_, err := interpretInputListMethods(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// Hard to predict exact order after shuffle; if no error occurs, assume pass
}

func TestInterpreter_ListLickFlatten(t *testing.T) {
	code := `
l -> [[1,2], [3,4]]
l.lick()
`
	result, err := interpretInputListMethods(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	list, ok := result.([]interface{})
	if !ok || len(list) != 4 {
		t.Errorf("Expected flattened list [1 2 3 4], got %v", result)
	}
}

func TestInterpreter_ListNestChunk(t *testing.T) {
	code := `
l -> [1,2,3,4,5]
l.nest(2)
`
	result, err := interpretInputListMethods(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	chunks, ok := result.([]interface{})
	if !ok || len(chunks) != 3 {
		t.Errorf("Expected 3 chunks, got %v", result)
	}
}

func TestInterpreter_ListHowlAtThreshold(t *testing.T) {
	code := `
l -> [1,5,10]
l.howl_at(4)
`
	result, err := interpretInputListMethods(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	filtered, ok := result.([]interface{})
	if !ok || len(filtered) != 2 {
		t.Errorf("Expected [5,10], got %v", result)
	}
}
