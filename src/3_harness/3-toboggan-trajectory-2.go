package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"github.com/calvinnhieu/advent-of-go/src/common"
)

// Approach:
// Read input line by line. For each line:
// - Starting at i == 0, check if tree (#) or open (.) at every (i + x % len) for (x E [1, 3, 5, 7]).
// - Count number of trees for each x.
// - for Down 2, skip row.

func main() {
	// Simple slopes to compute. Down 2 is not included.
	slopes := []int{1, 3, 5, 7}
	// Stores number of trees encountered.
	num_trees := make([]int, 5)
	// Stores index of current horizontal position.
	pos := make([]int, 5)
	// Stores number of lines processed.
	num_lines := 0
	

	// Open input file.
	_, src_path, _, _ := runtime.Caller(0)
	file, err := os.Open(common.Input_file(src_path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read input line by line.
	// For every slope, check tree (#) in current pos.
	// If tree, increment respective count in num_trees slice.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Cases: down slope = 1
		for i, slope := range slopes {
			if string(line[pos[i]]) == "#" {
				num_trees[i]++
			}
			pos[i] = (pos[i] + slope) % len(line)
		}

		// Special Case: down slope = 2
		if num_lines % 2 == 0 {
			if string(line[pos[4]]) == "#" {
				num_trees[4] = num_trees[4] + 1
			}
			pos[4] = (pos[4] + 1) % len(line)
		}
		num_lines++
	}
	
	// Handle scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("You will encounter %d, %d, %d, %d, and %d trees on your way down.\n", num_trees[0], num_trees[1], num_trees[2], num_trees[3], num_trees[4])
	fmt.Printf("The product is %d.\n", num_trees[0] * num_trees[1] * num_trees[2] * num_trees[3] * num_trees[4])
}
