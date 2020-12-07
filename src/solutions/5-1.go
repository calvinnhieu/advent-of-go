package solutions

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

// Takes in a boarding pass seat string and returns its binary number representation as a string.
func ToBinaryString(boardingPass string) string {
	var bString, char string

	// Parse each rune as 0 or 1.
	// Assumes input includes only F, B, L, R.
	for _, r := range boardingPass {
		char = string(r)

		if char == "F" || char == "L" {
			bString = bString + "0"
		} else if char == "B" || char == "R" {
			bString = bString + "1"
		}
	}

	return bString
}

// Converts a string representing a binary number into its decimal representation.
func ToDecimal(bString string) int64 {
	num, err := strconv.ParseInt(bString, 2, 10)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(bString, num)

	return num
}

// Solution for the Binary Boarding problem (5.1):
func BinaryBoardingA() (func(string), func()) {
	// Define state variables.
	var row, col int64
	var maxId float64

	// Processes line input and stores state in closure.
	processor := func(line string) {
		// Convert first 7 runes to binary number string, convert to decimal -> row
		row = ToDecimal(ToBinaryString(line[:7]))
		// Repeat for last 3 runes -> col
		col = ToDecimal(ToBinaryString(line[7:]))
		// Compute id = row * 8 + col
		maxId = math.Max(float64((row * 8) + col), maxId)
	}

	// Prints current closure state.
	result := func() {
		fmt.Println(maxId)
	}

	return processor, result
}
