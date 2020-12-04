// The common package provides shared utility functions for all Advent of Go solutions.
package common

import (
	"fmt"
	"os"
)

var ProblemName = map[int]string{
	1: "report-repair",
	2: "password-philosophy",
	3: "toboggan-trajectory",
}

// Returns the input file path given the day of the problem.
func Input_file(day int) string {
	input_dir := os.Getenv("AOG_INPUT_DIR")
	return fmt.Sprintf("%s/%d-%s.txt", input_dir, day, ProblemName[day])
}
