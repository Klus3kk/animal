package interpreter

import (
	"animal/core"
	"os"
	"testing"
)

func interpretInputBestiary(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_BestiaryImportVariables(t *testing.T) {
	bestiaryCode := `
x -> 42
y -> "animal"
!shelter -> ["x", "y"]
`
	// Write a temporary bestiary file
	fileName := "test_bestiary.anml"
	err := os.WriteFile(fileName, []byte(bestiaryCode), 0644)
	if err != nil {
		t.Fatalf("Failed to write bestiary file: %v", err)
	}
	defer os.Remove(fileName)

	code := `
%bestiary "test_bestiary.anml"
x meow 1 sniffback
`
	result, err := interpretInputBestiary(t, code)
	if err != nil {
		t.Fatalf("Unexpected error during bestiary import: %v", err)
	}

	if num, ok := result.(float64); !ok || num != 43 {
		t.Errorf("Expected imported x to be 42, plus 1 = 43, got %v", result)
	}
}

func TestInterpreter_BestiaryImportOnlyShelterSymbols(t *testing.T) {
	bestiaryCode := `
a -> 1
b -> 2
!shelter -> ["a"]
`
	fileName := "test_bestiary2.anml"
	err := os.WriteFile(fileName, []byte(bestiaryCode), 0644)
	if err != nil {
		t.Fatalf("Failed to write bestiary file: %v", err)
	}
	defer os.Remove(fileName)

	code := `
%bestiary "test_bestiary2.anml"
b
`
	_, err = interpretInputBestiary(t, code)
	if err == nil {
		t.Fatalf("Expected error accessing 'b' (not in shelter), got none")
	}
}
