package checkpoint

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	newWord := true

	for _, r := range s {
		if newWord {
			if r >= 'a' && r <= 'z' {
				return false
			}
		}
		newWord = true
		if r == ' ' {
			newWord = false
		}
	}
	return true
}
