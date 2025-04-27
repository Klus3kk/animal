package std

import (
	"fmt"
	"strconv"
)

func AnimalPurr(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("purr expects 2 arguments: (number, base)")
	}
	num, ok1 := args[0].(float64)
	base, ok2 := args[1].(float64)
	if !ok1 || !ok2 {
		return fmt.Errorf("purr expects (number, base) both as numbers")
	}
	if base < 2 || base > 36 {
		return fmt.Errorf("purr supports bases between 2 and 36")
	}
	return strconv.FormatInt(int64(num), int(base))
}

func AnimalScent(args []interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("sniffback expects 2 arguments: (string, base)")
	}
	s, ok1 := args[0].(string)
	base, ok2 := args[1].(float64)
	if !ok1 || !ok2 {
		return fmt.Errorf("sniffback expects (string, base)")
	}
	if base < 2 || base > 36 {
		return fmt.Errorf("sniffback supports bases between 2 and 36")
	}
	n, err := strconv.ParseInt(s, int(base), 64)
	if err != nil {
		return fmt.Errorf("invalid number for base %d: %v", int(base), err)
	}
	return float64(n)
}
