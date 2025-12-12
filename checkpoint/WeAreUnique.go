package checkpoint

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	str1Runes := make(map[rune]bool)
	str2Runes := make(map[rune]bool)

	for _, r := range str1 {
		str1Runes[r] = true
	}

	for _, r := range str2 {
		str2Runes[r] = true
	}

	count := 0
	for r := range str1Runes {
		if !str2Runes[r] {
			count++
		}
	}

	for r := range str2Runes {
		if !str1Runes[r] {
			count++
		}
	}
	return count
}
