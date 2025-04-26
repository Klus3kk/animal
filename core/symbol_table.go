package core

import "fmt"

// SYMBOL TABLE
type SymbolTable struct {
	symbols map[string]interface{} // Dictionary to store symbols
	parent  *SymbolTable           // Pointer to parent symbol table
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make(map[string]interface{}), // Initialize symbols map
		parent:  nil,
	}
}

// Set a value in the symbol table
func (s *SymbolTable) Set(name string, value interface{}) {
	if name == "" {
		fmt.Println("Error: Variable name cannot be empty")
		return
	}
	if Debug {
		fmt.Printf("Setting variable %s to %v in context\n", name, value)
	}
	s.symbols[name] = value
}

// Get a value from the symbol table
func (s *SymbolTable) get(name string) interface{} {
	if name == "" {
		fmt.Println("Error: Variable name cannot be empty")
		return nil
	}
	if value, exists := s.symbols[name]; exists {
		if Debug {
			fmt.Printf("Getting variable %s with value %v from context\n", name, value)
		}
		return value
	} else if s.parent != nil {
		if Debug {
			fmt.Printf("Looking up variable %s in parent context\n", name)
		}
		return s.parent.get(name)
	}
	fmt.Printf("Variable %s not found\n", name) // Added debug
	return nil
}

// Remove a symbol from the table
func (s *SymbolTable) remove(name string) {
	delete(s.symbols, name)
}
