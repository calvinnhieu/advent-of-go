package solutions

import (
	"fmt"
)

// Solution for the Toboggan Trajectory problem (3.1):
// For each consequtive line, check the next 3rd position (wraparound) for a tree.
// Count number of trees and return.
func TobogganTrajectoryA() (func(string), func()) {
	// Stores number of trees encountered.
	num_trees := 0
	// Stores index of current horizontal position
	i := 0

	// Processes line input and stores state in closure.
	// Check every 3rd element (with wraparound) in sequential rows for presence of tree (#).
	// If tree, increment num_trees.
	processor := func(line string) {
		if string(line[i]) == "#" {
			num_trees++
		}
		i = (i + 3) % len(line)
	}

	// Prints current closure state.
	result := func() {
		fmt.Printf("A: You will encounter %d trees on your way down.\n", num_trees)
	}

	return processor, result
}
