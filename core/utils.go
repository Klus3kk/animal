package core

// Utility function to check if a TokenType is in the list
func contains(ops []string, op string) bool {
	for _, val := range ops {
		if val == op {
			return true
		}
	}
	return false
}
