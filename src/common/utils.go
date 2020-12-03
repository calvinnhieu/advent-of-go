// The common package provides shared utility functions for all Advent of Go solutions.
package common

import (
	"fmt"
	"os"
	"strings"
)

// Returns the input file path given the absolute path of a solution file.
func Input_file(src_path string) string {
	input_dir := os.Getenv("AOG_INPUT_DIR")
	input_prefix := src_path[strings.LastIndex(src_path, "/") + 1:len(src_path) - 5]
	return fmt.Sprintf("%s/%s.txt", input_dir, input_prefix)
}
