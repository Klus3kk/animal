package std

import "animal/core"

// RegisterStandardLibrary injects standard library functions into the symbol table.
func RegisterStandardLibrary(symbolTable *core.SymbolTable) {
	// Math
	symbolTable.Set("max", AnimalMax)
	symbolTable.Set("min", AnimalMin)
	symbolTable.Set("abs", AnimalAbs)
	// Numbers
	symbolTable.Set("purr", AnimalPurr)
	symbolTable.Set("scent", AnimalScent)
	// Random
	symbolTable.Set("pounce", AnimalPounce)
	symbolTable.Set("stalk", AnimalStalk)
	symbolTable.Set("tumble", AnimalTumble)
	// List
	symbolTable.Set("paw", AnimalPaw)
	symbolTable.Set("nuzzle", AnimalNuzzle)
	symbolTable.Set("burrow", AnimalBurrow)
	symbolTable.Set("perch", AnimalPerch)
	symbolTable.Set("lick", AnimalLick)
	symbolTable.Set("howl", AnimalHowl)
	symbolTable.Set("chase", AnimalChase)
	symbolTable.Set("trace", AnimalTrace)
	symbolTable.Set("trail", AnimalTrail)
	symbolTable.Set("pelt", AnimalPelt)
	symbolTable.Set("howlpack", AnimalHowlpack)
	symbolTable.Set("nest", AnimalNest)

}
