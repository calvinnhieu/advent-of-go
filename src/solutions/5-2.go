package solutions

import (
	"fmt"
	"sort"
)

// Solution for the Binary Boarding problem (5.2):
func BinaryBoardingB() (func(string), func()) {
	// Define state variables.
	var row, col int64
	var ids []int

	// Processes line input and stores state in closure.
	processor := func(line string) {
		// Convert first 7 runes to binary number string, convert to decimal -> row
		row = ToDecimal(ToBinaryString(line[:7]))
		// Repeat for last 3 runes -> col
		col = ToDecimal(ToBinaryString(line[7:]))
		// Compute id = row * 8 + col
		ids = append(ids, int(row * 8 + col))
	}

	// Prints current closure state.
	result := func() {
		sort.Ints(ids)
		first := ids[0]

		// Finds the first non-consequtive ID.
		for i, id := range ids {
			if id != first + i {
				fmt.Println(first + i)
				return
			}
		}
	}

	return processor, result
}
