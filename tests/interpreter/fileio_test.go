package interpreter

import (
	"animal/core"
	"os"
	"testing"
)

func interpretInputFileIO(t *testing.T, input string) (interface{}, error) {
	t.Helper()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<program>",
		Symbol_Table: globalSymbolTable,
	}

	return core.CustomRun(input, "<stdin>", context)
}

func TestInterpreter_DropAndFetchFile(t *testing.T) {
	fileName := "testfile.txt"

	code := `
drop("testfile.txt", "hello world")
fetch("testfile.txt")
`
	result, err := interpretInputFileIO(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer os.Remove(fileName)

	if str, ok := result.(string); !ok || str != "hello world" {
		t.Errorf("Expected file content 'hello world', got %v", result)
	}
}

func TestInterpreter_DropAppendFile(t *testing.T) {
	fileName := "appendfile.txt"

	code := `
drop("appendfile.txt", "hello")
drop_append("appendfile.txt", " world")
fetch("appendfile.txt")
`
	result, err := interpretInputFileIO(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer os.Remove(fileName)

	if str, ok := result.(string); !ok || str != "hello world" {
		t.Errorf("Expected file content 'hello world', got %v", result)
	}
}

func TestInterpreter_SniffFileExists(t *testing.T) {
	fileName := "snifffile.txt"
	os.WriteFile(fileName, []byte("data"), 0644)
	defer os.Remove(fileName)

	code := `
sniff_file("snifffile.txt")
`
	result, err := interpretInputFileIO(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if exists, ok := result.(bool); !ok || !exists {
		t.Errorf("Expected file to exist, got %v", result)
	}
}

func TestInterpreter_SniffFileDoesNotExist(t *testing.T) {
	code := `
sniff_file("nonexistent.txt")
`
	result, err := interpretInputFileIO(t, code)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if exists, ok := result.(bool); !ok || exists {
		t.Errorf("Expected file not to exist, got %v", result)
	}
}
