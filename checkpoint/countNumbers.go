package checkpoint

func CountNumbers(arg string) int {
	count := 0
	for _, char := range arg {
		if char >= '0' && char <= '9' {
			count++
		}
	}
	return count
}
