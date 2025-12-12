package checkpoint

func DigitLen(n, base int) int {
	if base > 36 || base < 2 {
		return -1
	}

	if n < 0 {
		n = -n
	}
	ans := 0

	for n > 0 {
		n = n / base
		ans++
	}
	return ans
}
