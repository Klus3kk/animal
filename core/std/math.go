package std

import (
	"fmt"
)

// AnimalMax returns the maximum of two numbers.
func AnimalMax(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("max expects 2 arguments")
	}
	a, ok1 := args[0].(float64)
	b, ok2 := args[1].(float64)
	if !ok1 || !ok2 {
		return fmt.Errorf("max expects numbers")
	}
	if a > b {
		return a
	}
	return b
}

// AnimalMin returns the minimum of two numbers.
func AnimalMin(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("min expects 2 arguments")
	}
	a, ok1 := args[0].(float64)
	b, ok2 := args[1].(float64)
	if !ok1 || !ok2 {
		return fmt.Errorf("min expects numbers")
	}
	if a < b {
		return a
	}
	return b
}

// AnimalAbs returns the absolute value of a number.
func AnimalAbs(args []interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("abs expects 1 argument")
	}
	a, ok := args[0].(float64)
	if !ok {
		return fmt.Errorf("abs expects a number")
	}
	if a < 0 {
		return -a
	}
	return a
}
