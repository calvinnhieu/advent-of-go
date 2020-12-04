package solutions

import (
	"fmt"
)

// Solution for the Example problem (0.0):
func ExampleA() (func(string), func()) {
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
