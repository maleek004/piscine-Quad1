package checkpoint

import "fmt"

func Addprimesum(n int) {
	if n < 0 {
		fmt.Println("0")
		return
	}
	res := 0
	for i := n; i >= 2; i-- {
		if IsPrime(i) {
			res += i
		}
	}
	fmt.Println(res)
}
