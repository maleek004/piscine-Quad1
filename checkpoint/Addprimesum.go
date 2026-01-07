package checkpoint

import (
	"os"

	"github.com/01-edu/z01"
)

// the main logic that we learned here is how to convert an integer into a list of numbers
// and how to convert a list of numbers back to an integer
func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := make([]rune, 0)
	for n > 0 {
		digits = append([]rune{rune(n%10) + '0'}, digits...)
		n /= 10
	}

	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}

func AddPrimeSum(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			sum += i
		}
	}
	return sum
}

// this function is to be used in the main program
func AddPrimeSumMain() {
	// if we have no or more than one input, or our input is not a number ,
	// we should print 0 using our PrintInt function which relies on only PrintRune
	if len(os.Args) != 2 || os.Args[1][0] < '0' || os.Args[1][0] > '9' {
		PrintInt(0)
		z01.PrintRune('\n')
		return
	}

	// constructing an integer from the 'string number' input, by starting with a base int of 0
	// and for each character in the 'string number', multiplying our base by 10 and adding the charcater of the 'string number' to it
	// after the character has been converted into a proper integer
	n := 0
	for _, c := range os.Args[1] {
		if c < '0' || c > '9' { // if we find any non number , we are quitting the program
			PrintInt(0)
			z01.PrintRune('\n')
			return
		}
		n = n*10 + int(c-'0')
	}

	if n <= 0 {
		PrintInt(0)
		z01.PrintRune('\n')
		return
	}

	result := AddPrimeSum(n)
	PrintInt(result)
	z01.PrintRune('\n')
}
