package main

import (
	"animal/core"
	"syscall/js"
)

func runAnimal(this js.Value, args []js.Value) interface{} {
	code := args[0].String()

	globalSymbolTable := core.NewSymbolTable()
	context := &core.Context{
		DisplayName:  "<browser>",
		Symbol_Table: globalSymbolTable,
	}

	result, err := core.CustomRun(code, "<stdin>", context)
	if err != nil {
		return js.ValueOf("Error: " + err.Error())
	}
	if result != nil {
		return js.ValueOf(result)
	}
	return js.ValueOf("Finished with no output")
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("runAnimalCode", js.FuncOf(runAnimal))
	<-c // block forever
}
