package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/calvinnhieu/advent-of-go/src/common"
)

// Reads CLI args and returns the following args:
// - day int
func parseArgs() int {
	args := os.Args[1:]
	
	// Except day arg.
	if len(args) < 1 {
		log.Fatal("Please specify a day (1-25) as an input parameter.")
	}

	// Check day type (int).
	day, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Day must be int (1-25).")
	}

	// Validate day value.
	if _, ok := common.SolutionMap[day]; !ok {
		log.Fatal(fmt.Sprintf("There is no solution for day %d.\n", day))
	}

	return day
}

func main() {
	// Read CLI args.
	day := parseArgs()

	// Dynamically load solution according to day.	
	processorA, resultA := common.SolutionMap[day][0]()
	processorB, resultB := common.SolutionMap[day][1]()

	// Open input file.
	file, err := os.Open(common.Input_file(day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read input line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		processorA(line)
		processorB(line)
	}
	
	// Handle scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	resultA()
	resultB()
}
