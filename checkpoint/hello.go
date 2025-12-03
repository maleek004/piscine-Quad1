package checkpoint

import "github.com/01-edu/z01"

func Hello() {
	sentence := "Hello World!"
	for _, char := range sentence {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
