package tests

import (
	"animal/core"
	"testing"
)

func TestSymbolTable_SetAndGetSymbol(t *testing.T) {
	table := core.NewSymbolTable()
	table.Set("x", 10)

	symbol, exists := table.Get("x")
	if !exists {
		t.Fatalf("Expected symbol 'x' to exist")
	}

	if value, ok := symbol.Value.(int); !ok || value != 10 {
		t.Errorf("Expected 'x' = 10, got %v", symbol.Value)
	}
}

func TestSymbolTable_TypeEnforcedSet(t *testing.T) {
	table := core.NewSymbolTable()
	table.SetWithType("x", 42, "int")

	symbol, exists := table.Get("x")
	if !exists {
		t.Fatalf("Expected symbol 'x' to exist")
	}

	if symbol.Type != "int" {
		t.Errorf("Expected type 'int', got %v", symbol.Type)
	}
}

func TestSymbolTable_ParentLookup(t *testing.T) {
	parent := core.NewSymbolTable()
	parent.Set("parent_var", 100)

	child := core.NewSymbolTable()
	child.SetParent(parent)

	symbol, exists := child.Get("parent_var")
	if !exists {
		t.Fatalf("Expected to find 'parent_var' via parent")
	}

	if value, ok := symbol.Value.(int); !ok || value != 100 {
		t.Errorf("Expected 'parent_var' = 100, got %v", symbol.Value)
	}
}

func TestSymbolTable_RemoveSymbol(t *testing.T) {
	table := core.NewSymbolTable()
	table.Set("temp", 5)

	table.Remove("temp")

	_, exists := table.Get("temp")
	if exists {
		t.Errorf("Expected 'temp' to be removed")
	}
}
