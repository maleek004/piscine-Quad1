package checkpoint

func CountAlpha(arg string) int {
	count := 0
	for _, char := range arg {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
	}
	return count
}
