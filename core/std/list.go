package std

import (
	"fmt"
)

// paw(x, min, max): clamp a number
func AnimalPaw(args []interface{}) interface{} {
	if len(args) != 3 {
		return fmt.Errorf("paw expects 3 arguments: (value, min, max)")
	}
	val, ok1 := args[0].(float64)
	min, ok2 := args[1].(float64)
	max, ok3 := args[2].(float64)
	if !ok1 || !ok2 || !ok3 {
		return fmt.Errorf("paw expects numbers")
	}
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

// nuzzle(a, b): merge two lists or two strings
func AnimalNuzzle(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("nuzzle expects 2 arguments")
	}

	switch a := args[0].(type) {
	case []interface{}:
		if b, ok := args[1].([]interface{}); ok {
			return append(a, b...)
		}
	case string:
		if b, ok := args[1].(string); ok {
			return a + b
		}
	}

	return fmt.Errorf("nuzzle expects (list,list) or (string,string)")
}

// burrow(n): create list of n nil elements
func AnimalBurrow(args []interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("burrow expects 1 argument")
	}
	size, ok := args[0].(float64)
	if !ok {
		return fmt.Errorf("burrow expects a number")
	}
	if size < 0 {
		return fmt.Errorf("burrow size must be non-negative")
	}
	list := make([]interface{}, int(size))
	for i := 0; i < int(size); i++ {
		list[i] = nil
	}
	return list
}

// perch(list): return all permutations of a list
func AnimalPerch(args []interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("perch expects 1 argument")
	}
	list, ok := args[0].([]interface{})
	if !ok {
		return fmt.Errorf("perch expects a list")
	}

	return permute(list)
}

// Helper: generate all permutations
func permute(list []interface{}) []interface{} {
	var res []interface{}
	var helper func([]interface{}, int)

	helper = func(a []interface{}, n int) {
		if n == 1 {
			tmp := make([]interface{}, len(a))
			copy(tmp, a)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(a, n-1)
				if n%2 == 1 {
					a[0], a[n-1] = a[n-1], a[0]
				} else {
					a[i], a[n-1] = a[n-1], a[i]
				}
			}
		}
	}

	cpy := make([]interface{}, len(list))
	copy(cpy, list)
	helper(cpy, len(cpy))
	return res
}

// lick(list): flatten nested list into single list
func AnimalLick(args []interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("lick expects 1 argument (list)")
	}
	list, ok := args[0].([]interface{})
	if !ok {
		return fmt.Errorf("lick expects a list")
	}
	return flatten(list)
}

// Helper: recursively flatten list
func flatten(list []interface{}) []interface{} {
	var flat []interface{}
	for _, item := range list {
		if sublist, ok := item.([]interface{}); ok {
			flat = append(flat, flatten(sublist)...)
		} else {
			flat = append(flat, item)
		}
	}
	return flat
}

// howl(list, item): find index of item inside list
func AnimalHowl(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("howl expects 2 arguments (list, item)")
	}
	list, ok := args[0].([]interface{})
	if !ok {
		return fmt.Errorf("howl expects a list as first argument")
	}
	item := args[1]

	for idx, v := range list {
		if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", item) {
			return float64(idx)
		}
	}
	return float64(-1) // not found
}

// chase(x, n): repeat element x n times into a list
func AnimalChase(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("chase expects 2 arguments (element, times)")
	}
	element := args[0]
	n, ok := args[1].(float64)
	if !ok {
		return fmt.Errorf("chase expects second argument as number")
	}
	if n < 0 {
		return fmt.Errorf("chase count must be non-negative")
	}

	result := []interface{}{}
	for i := 0; i < int(n); i++ {
		result = append(result, element)
	}
	return result
}

func AnimalTrace(args []interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("trace expects 1 argument (list)")
	}
	list, ok := args[0].([]interface{})
	if !ok {
		return fmt.Errorf("trace expects a list")
	}

	result := []interface{}{}
	var runningTotal float64 = 0

	for _, v := range list {
		num, ok := v.(float64)
		if !ok {
			return fmt.Errorf("trace expects a list of numbers")
		}
		runningTotal += num
		result = append(result, runningTotal)
	}
	return result
}

func AnimalTrail(args []interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("trail expects 1 argument (list)")
	}
	list, ok := args[0].([]interface{})
	if !ok {
		return fmt.Errorf("trail expects a list")
	}

	result := []interface{}{}
	for i := 1; i <= len(list); i++ {
		sublist := make([]interface{}, i)
		copy(sublist, list[:i])
		result = append(result, sublist)
	}
	return result
}

func AnimalPelt(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("pelt expects 2 arguments (value, times)")
	}
	n := args[0]
	times, ok := args[1].(float64)
	if !ok {
		return fmt.Errorf("pelt expects (value, times)")
	}

	str := fmt.Sprintf("%v", n)
	repeated := ""
	for i := 0; i < int(times); i++ {
		repeated += str
	}
	return repeated
}

func AnimalHowlpack(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("howlpack expects 2 arguments (list, item)")
	}
	list, ok := args[0].([]interface{})
	if !ok {
		return fmt.Errorf("howlpack expects a list as first argument")
	}
	target := args[1]

	indexes := []interface{}{}
	for idx, v := range list {
		if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", target) {
			indexes = append(indexes, float64(idx))
		}
	}
	return indexes
}

func AnimalNest(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("nest expects 2 arguments (value, depth)")
	}
	value := args[0]
	depth, ok := args[1].(float64)
	if !ok {
		return fmt.Errorf("nest expects (value, depth)")
	}
	if depth < 0 {
		return fmt.Errorf("nest depth must be non-negative")
	}

	result := value
	for i := 0; i < int(depth); i++ {
		result = []interface{}{result}
	}
	return result
}
