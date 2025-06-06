package core

import "fmt"

// SYMBOL and SYMBOL TABLE

type Symbol struct {
	Value interface{}
	Type  string // "" means dynamic (no type enforced)
}

type SymbolTable struct {
	symbols map[string]Symbol
	parent  *SymbolTable
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make(map[string]Symbol),
		parent:  nil,
	}
}

// Set a value dynamically (no type enforced)
func (s *SymbolTable) Set(name string, value interface{}) {
	if name == "" {
		fmt.Println("Error: Variable name cannot be empty")
		return
	}
	s.symbols[name] = Symbol{Value: value, Type: ""}
}

// Set a value with a given type
func (s *SymbolTable) SetWithType(name string, value interface{}, typeName string) {
	if name == "" {
		fmt.Println("Error: Variable name cannot be empty")
		return
	}
	s.symbols[name] = Symbol{Value: value, Type: typeName}
}

// Get retrieves the symbol (value + type)
func (s *SymbolTable) Get(name string) (Symbol, bool) {
	if name == "" {
		fmt.Println("Error: Variable name cannot be empty")
		return Symbol{}, false
	}
	if value, exists := s.symbols[name]; exists {
		return value, true
	}
	if s.parent != nil && s.parent != s {
		return s.parent.Get(name)
	}
	return Symbol{}, false
}

// Remove deletes a symbol
func (s *SymbolTable) Remove(name string) {
	delete(s.symbols, name)
}

func (s *SymbolTable) SetParent(parent *SymbolTable) {
	s.parent = parent
}
