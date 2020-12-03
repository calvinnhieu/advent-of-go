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

func main() {
	// Open input file.
	_, src_path, _, _ := runtime.Caller(0)
	file, err := os.Open(common.Input_file(src_path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read template input contents.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	
	// Handle scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
