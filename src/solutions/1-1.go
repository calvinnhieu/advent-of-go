package solutions

import (
	"fmt"
	"log"
	"strconv"
)

// Solution for the Report Repair problem (1.1):
func ReportRepairA() (func(string), func()) {
	// Two-sum target value.
	target := 2020
	// Counts occurrences of each number.
	// Assumes numbers are not unique
	numCounts := make(map[int]int)
	// Stores two-sum values.
	num1, num2 := 0, 0

	// Processes line input and stores state in closure.
	processor := func(line string) {
		// Parse line to number
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		// If the current number's complement is in the map,
		// then we have found our 2-sum and yield the product.
		if _, ok := numCounts [target - num]; ok {
			num1, num2 = num, target - num
			return
		}

		// Insert number into count map
		if _, ok := numCounts[num]; !ok {
			numCounts[num] = 0
		}
		numCounts[num] = numCounts[num] + 1
	}

	// Prints current closure state.
	result := func() {
		fmt.Printf("The product of the 2-sum is: %d * %d = %d\n", num1, num2, num1 * num2)
	}

	return processor, result
}
