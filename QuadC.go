package main

import "fmt"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	countMidCells := x - 2
	printToprow := func() {
		fmt.Print("A")
		if countMidCells >= 1 {
			for i := countMidCells; i > 0; i-- {
				fmt.Print("B")
			}
			fmt.Println("A")
		} else {
			fmt.Print("\n")
		}
	}
	printMiddleRow := func() {
		fmt.Print("B")
		if countMidCells >= 1 {
			for i := countMidCells; i > 0; i-- {
				fmt.Print(" ")
			}
			fmt.Println("B")
		} else {
			fmt.Print("\n")
		}
	}
	printButtomRow := func() {
		fmt.Print("C")
		if countMidCells >= 1 {
			for i := countMidCells; i > 0; i-- {
				fmt.Print("B")
			}
			fmt.Println("C")
		} else {
			fmt.Print("\n")
		}
	}
	// Now use the helper function to pront the box
	printToprow()
	for i := y - 2; i >= 1; i-- {
		printMiddleRow()
	}
	if y > 1 {
		printButtomRow()
	}
}
