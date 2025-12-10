package main

import (
	"fmt"
	"os"
	"piscine/Quad"
	"piscine/checkpoint"
	"piscine/piscine"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	Quad.QuadA(5, 3)
	fmt.Println()
	Quad.QuadB(5, 3)
	fmt.Println()
	Quad.QuadC(5, 3)
	fmt.Println()
	Quad.QuadD(5, 3)
	fmt.Println()
	Quad.QuadE(5, 3)
	fmt.Println()

	// testing the First Rune function
	fmt.Println(strings.Repeat("=", 10))
	firstWord := "Maleek"
	secondWord := "Ola"
	fmt.Println("testing the FirstRune function with strings '", firstWord, "' and '", secondWord, "'")
	fmt.Println(string(piscine.FirstRune(firstWord))) // using the builtin string() function to convert the rune back to string
	z01.PrintRune(piscine.FirstRune(secondWord))      // using z01's PrintRune function to print the rune as string directly
	fmt.Println()
	fmt.Println("for the sake of clarity , the int32 runes for these words are :")
	fmt.Println([]rune(firstWord))
	fmt.Println([]rune(secondWord))

	// testing the LastRune function
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("testing the LastRune function with strings '", firstWord, "' and '", secondWord, "'")
	fmt.Println(string(piscine.LastRune(firstWord))) // using the builtin string() function to convert the rune back to string
	z01.PrintRune(piscine.LastRune(secondWord))      // using z01's PrintRune function to print the rune as string directly
	fmt.Println()

	// testing the NRune function
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("testing the NRune function by returning the rune at index 2 for strings '", firstWord, "' and '", secondWord, "'")
	fmt.Println(string(piscine.NRune(firstWord, 2)))
	fmt.Println(string(piscine.NRune(secondWord, 2)))

	// testing the CheckNumber function
	fmt.Println(strings.Repeat("=", 10))
	githubUserName := "maleek004"
	fmt.Println("testing the CheckNumber function (from checkpoints) with the string'", githubUserName, "' ")
	isNumber := checkpoint.CheckNumber(githubUserName)
	fmt.Println("Is number:", isNumber)

	// testing the CountAlpha function
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("Testing the CountAlpha function (from checkpoint) with the string '", githubUserName, "'")
	result := checkpoint.CountAlpha(githubUserName)
	fmt.Println("Result:", result)

	// testing the CountNumbers function
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("Testing the CountNumbers function (from checkpoint) with the string '", githubUserName, "'")
	CountNumbers_result := checkpoint.CountNumbers(githubUserName)
	fmt.Println("Result:", CountNumbers_result)

	// testing the CountChar function
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("Testing the CountChar function (from checkpoint) by counting occurences of 'e' in the string '", githubUserName, "'")
	CountChar_result := checkpoint.CountChar(githubUserName, 'e')
	fmt.Println("Result:", CountChar_result)

	// testing the CamelToSnakeCase function
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("Testing the CamelToSnakeCase function (from checkpoint) ")
	fmt.Println(checkpoint.CamelToSnakeCase("HelloWorld"))
	fmt.Println(checkpoint.CamelToSnakeCase("helloWorld"))
	fmt.Println(checkpoint.CamelToSnakeCase("camelCase"))
	fmt.Println(checkpoint.CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(checkpoint.CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(checkpoint.CamelToSnakeCase("hey2"))
	fmt.Println("----------------------- I still need to work on the camelcase code")

	// testing the FoodDeliveryTime function
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("Testing the FoodDeliveryTime function (from checkpoint) ")
	fmt.Println(checkpoint.FoodDeliveryTime("burger"))
	fmt.Println(checkpoint.FoodDeliveryTime("chips"))
	fmt.Println(checkpoint.FoodDeliveryTime("nuggets"))
	fmt.Println(checkpoint.FoodDeliveryTime("burger") + checkpoint.FoodDeliveryTime("chips") + checkpoint.FoodDeliveryTime("nuggets"))

	
	// testing the Gcd function
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("Testing the Gcd function (from checkpoint) ")
	fmt.Println(checkpoint.Gcd(42, 10))
	fmt.Println(checkpoint.Gcd(42, 12))
	fmt.Println(checkpoint.Gcd(14, 77))
	fmt.Println(checkpoint.Gcd(17, 3))
	
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return // Stop the program
	}

	// STEP 2: Create an empty 9×9 grid
	var grid [9][9]int
	// This makes:
	// [0][0][0][0][0][0][0][0][0]
	// [0][0][0][0][0][0][0][0][0]
	// ... (9 rows of 9 zeros)

	// STEP 3: Fill the grid with what user typed
	for i := 0; i < 9; i++ {
		// Get one row from user's input
		row := os.Args[i+1] // +1 because Args[0] is program name

		// STEP 4: Make sure row is 9 characters long
		if len(row) != 9 {
			fmt.Println("Error")
			return
		}

		// STEP 5: Convert each character to a number
		for j := 0; j < 9; j++ {
			if row[j] == '.' {
				// Dot means empty cell
				grid[i][j] = 0
			} else if row[j] >= '1' && row[j] <= '9' {
				// Convert character '1'-'9' to number 1-9
				// Example: '5' - '0' = 5
				grid[i][j] = int(row[j] - '0')
			} else {
				// Invalid character (not . or 1-9)
				fmt.Println("Error")
				return
			}
		}
	}

	// STEP 6: Try to solve the puzzle!
	if Quad.SolveSudoku(&grid) {
		// Success! Print the solution
		Quad.PrintSudoku(grid)
	} else {
		// Impossible to solve
		fmt.Println("Error")
	}
}
