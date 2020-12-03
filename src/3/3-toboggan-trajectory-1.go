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
// - Starting at i == 0, check if tree (#) or open (.) at i + 3
// - Count number of trees
// - If i + 3 exceeds line length, take (i + 3 % length)
// -- More performant to mod everytime (skips check)

func main() {
	// Stores number of trees encountered.
	num_trees := 0
	// Stores index of current horizontal position
	i := 0

	// Open input file.
	_, src_path, _, _ := runtime.Caller(0)
	file, err := os.Open(common.Input_file(src_path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read input line by line.
	// Check every 3rd element, with wraparound, in sequential rows for tree (#).
	// If tree, count.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if string(line[i]) == "#" {
			num_trees++
		}
		i = (i + 3) % len(line)
	}
	
	// Handle scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You will encounter %d trees on your way down.\n", num_trees)
}
