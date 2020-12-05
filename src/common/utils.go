// The common package provides shared utility functions for all Advent of Go solutions.
package common

import (
	"fmt"
	"os"
	"github.com/calvinnhieu/advent-of-go/src/solutions"
)

// TODO: create Solution type
var SolutionMap = map[int][]func()(func(string), func()){
	1: []func()(func(string), func()){solutions.ReportRepairA, solutions.ReportRepairB},
	2: []func()(func(string), func()){solutions.PasswordPhilosophyA, solutions.PasswordPhilosophyB},
	3: []func()(func(string), func()){solutions.TobogganTrajectoryA, solutions.TobogganTrajectoryB},
	4: []func()(func(string), func()){solutions.PassportProcessingA, solutions.PassportProcessingB},
	5: []func()(func(string), func()){solutions.BinaryBoardingA, solutions.BinaryBoardingB},
}

var ProblemName = map[int]string{
	1: "report-repair",
	2: "password-philosophy",
	3: "toboggan-trajectory",
	4: "passport-processing",
	5: "binary-boarding",
}

// Returns the input file path given the day of the problem.
func Input_file(day int) string {
	input_dir := os.Getenv("AOG_INPUT_DIR")
	return fmt.Sprintf("%s/%d-%s.txt", input_dir, day, ProblemName[day])
}

