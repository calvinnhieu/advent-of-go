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
// For each password, determine validity: [O(n*m); n = lines, m = chars/line].
// - Iterate chars - store min, max, char, iterate to password start index.
// - Count char instances in password.
// - If count within min/max boundary, valid. Else, invalid. Count valids.

// Returns the input file path.
func input_file() string {
	input_dir := "../input"
	p_num := 2
	name := "password-philosophy"
	input_fmt := "txt"

	return fmt.Sprintf("%s/%d-%s.%s", input_dir, p_num, name, input_fmt)
}

// Parses and returns the following tokens from a password rule string:
// min int - minimum occurrence boundary.
// max int - maximum occurrence boundary.
// char_rule string - char specified by rule.
// str_i int - the index of the start of the password.
func tokenize(input string) (int, int, string, int) {
	min := -1
	max := -1
	char_rule := ""
	str_i := -1

	prev_int := false
	for i, char := range input {
		// Dynamically read integers of unknown length from the input string.
		// This code reads two integers and stores them in min and max.
		// TODO: refactor to a parse number function. (Done in part 2).
		if min == -1 || max == -1 || prev_int {
			if num, err := strconv.Atoi(string(char)); err == nil {
				//fmt.Printf("num: %d min: %d max: %d.\n", num, min, max)
				if min == -1 {
					min = num
				} else if max == -1 && prev_int {
					min = (min * 10) + num
				} else if !prev_int {
					max = num	
				} else {
					max = (max * 10) + num
				}
				prev_int = true
			} else {
				// fmt.Printf("%s is not an int.\n", string(char))
				prev_int = false
			}
		} else {
			prev_int = false
		}

		// Lookahead 1 to find ":" to find char rule and extrapolate str_i.
		// Tokenization complete.
		// Assume ":" appears in every line preceding a non-empty password.
		if string(input[i + 1]) == ":" {
			char_rule = string(char)
			str_i = i + 3
			break
		}
	}

	return min, max, char_rule, str_i
}

func main() {
	// For logging purposes.
	input_size := 0
	// Counts the number of valid passwords.
	valid_count := 0

	// Open input file.
	_, src_path, _, _ := runtime.Caller(0)
	file, err := os.Open(common.Input_file(src_path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read input file line by line.
	// Tokenize each line to retrieve min, max, char_rule, and str_i.
	// Beginning at str_i, iterate through the password to count char instances.
	// Determine validity using min and max bounds. Increment if valid.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input_size++
		line := string(scanner.Text())
		// fmt.Println(line)
		min, max, char_rule, str_i := tokenize(line)

		// fmt.Printf("At least %d %s's and at most %d %s's in %s.\n", min, char_rule, max, char_rule, string(line[str_i:]))

		char_count := 0
		for _, char := range line[str_i:] {
			if string(char) == char_rule {
				char_count++
			}
		}

		// fmt.Printf("Found %d %s's.", char_count, char_rule)

		if char_count >= min && char_count <= max {
			valid_count++
			// fmt.Println(" Valid.\n")
		} else {
			// fmt.Println(" Invalid.\n")
		}
	}

	fmt.Printf("There are %d/%d valid passwords.\n", valid_count, input_size)
}
