package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"github.com/calvinnhieu/advent-of-go/src/common"
	"github.com/calvinnhieu/advent-of-go/src/solutions" // TODO: create this package
)

func main() {
	// TODO: Accept command line parameters:
	//	 - Problem #
	// 
	// Given a problem #, the runner will:
	// - Identify, open, and begin reading input file.
	// - Call #.1 and #.2 solutions (closure functions) with every line.
	// - After all lines have been processed, somehow retrieve results.

	// Open input file.
	_, src_path, _, _ := runtime.Caller(0)
	file, err := os.Open(common.Input_file(src_path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read input line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// TODO: implement
		// solutions.LineProcessor[PROBLEM_NUM][0](scanner.Text())
		// solutions.LineProcessor[PROBLEM_NUM][1](scanner.Text())
	}
	
	// Handle scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// TODO: implement
	// solutions.Result[PROBLEM_NUM][0]()
	// solutions.Result[PROBLEM_NUM][1]()
}
