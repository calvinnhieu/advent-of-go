package solutions

import (
	"fmt"
)

// Solution for the PassportProcessing problem (4.1):
func PassportProcessingA() (func(string), func()) {
	// Init state variables.
	counter := 0

	// Processes line input and stores state in closure.
	processor := func(line string) {
		counter++
	}

	// Prints current closure state.
	result := func() {
		fmt.Println(counter)
	}

	return processor, result
}
