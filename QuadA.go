package main

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	printXAxis := func() {
		if x <= 2 {
			for i := x; i >= 1; i-- {
				fmt.Print("o")
			}
			fmt.Print("\n")
		} else {
			fmt.Print("o")
			for i := x - 2; i >= 1; i-- {
				fmt.Print("-")
			}
			fmt.Print("o\n")
		}
	}

	printYPipes := func() {
		if y > 2 {
			fmt.Print("|")
			spaces := x - 2
			if spaces > 0 {
				for i := spaces; i >= 1; i-- {
					fmt.Print(" ")
				}
				fmt.Print("|\n")
			} else {
				fmt.Print("\n")
			}
		}
	}
	// Now use the helper functions to print the box
	printXAxis()
	for i := y - 2; i >= 1; i-- {
		printYPipes()
	}
	if y > 1 {
		printXAxis()
	}

}
