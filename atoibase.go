package piscine

func AtoiBase(s string, base string) int {
	// Validate base
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)

	// Create a lookup map: rune → value
	val := make(map[rune]int)
	for i, r := range base {
		val[r] = i
	}

	// Convert
	result := 0
	for _, r := range s {
		result = result*baseLen + val[r]
	}

	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}

	return true
}
