package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"runtime"
	"github.com/calvinnhieu/advent-of-go/src/common"
)

// Approach:
// Read input file
// Per line:
// - Parse line into map {num: count}
// - Lookup (2020 - num)
// - Return product if found

func main() {
	// Two-sum target value
	target := 2020
	// Map to count number occurrences
	// Assume numbers are not unique
	num_counts := make(map[int]int)
	
	// Open input file.
	_, src_path, _, _ := runtime.Caller(0)
	file, err := os.Open(common.Input_file(src_path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse line to number
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		
		// If the current number's complement is in the map,
		// then we have found our 2-sum and yield the product.
		if _, ok := num_counts[target - num]; ok {
			fmt.Printf("2-sum found. Product: %d * %d = %d\n", num, target - num, (target - num) * num)
			return
		}

		// Insert number into count map
		if _, ok := num_counts[num]; !ok {
			num_counts[num] = 0
		}
		num_counts[num] = num_counts[num] + 1
	}
	
	// Handle scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
