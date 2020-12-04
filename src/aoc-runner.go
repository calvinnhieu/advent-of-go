package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"github.com/calvinnhieu/advent-of-go/src/common"
	"github.com/calvinnhieu/advent-of-go/src/solutions"
)

func main() {
	// Read CLI parameters for day input.
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Please specify a day (1-25) as an input parameter.")
	}
	// Int type check.
	day, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Day must be int (1-25).")
	}
	// TODO: add check to ensure solution exists for day
	// TODO: retrieve correct solution given day
	processorA, resultA := solutions.TobogganTrajectoryA()
	processorB, resultB := solutions.TobogganTrajectoryB()

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
