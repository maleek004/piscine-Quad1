package checkpoint

func CountChar(arg string, c rune) int {
	count := 0
	for _, char := range arg {
		if char == c {
			count++
		}
	}
	return count
}
