package solutions

import (
	"fmt"
)

// Solution for the Toboggan Trajectory problem (3.2):
// For each consequtive line, check the next position (wraparound) for every slope for a tree.
// Count number of trees, multiply, and return.
func TobogganTrajectoryB() (func(string), func()) {
	// Stores number of trees encountered per slope.
	var num_trees [5]int
	// Stores index of current horizontal position.
	var pos [5]int
	// Slopes to compute, excluding Down 2.
	slopes := [...]int{1, 3, 5, 7}
	// Stores the number of lines processed.
	num_lines := 0

	// Processes line input and stores state in closure.
	processor := func(line string) {
		// Cases: down 1
		for i, slope := range slopes {
			if string(line[pos[i]]) == "#" {
				num_trees[i]++
			}
			pos[i] = (pos[i] + slope) % len(line)
		}

		// Case: down 2
		if num_lines % 2 == 0 {
			if string(line[pos[4]]) == "#" {
				num_trees[4] = num_trees[4] + 1
			}
			pos[4] = (pos[4] + 1) % len(line)
		}
		num_lines++
	}

	// Prints current closure state.
	result := func() {
		fmt.Printf("B: You will encounter %d, %d, %d, %d, and %d trees on your way down.", num_trees[0], num_trees[1], num_trees[2], num_trees[3], num_trees[4])
		fmt.Printf(" The product is %d.\n", num_trees[0] * num_trees[1] * num_trees[2] * num_trees[3] * num_trees[4])
	}

	return processor, result
}
