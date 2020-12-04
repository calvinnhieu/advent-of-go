package solutions

import (
	"fmt"
	"strconv"
)

// Reads a number token from a string and advances the given index past the end of the number.
// Returns the parsed number and new index.
// Assumes the starting index is a number, else returns 0.
// Assumes the read number does not reach the end of input, else fails.
func readNum(input string, curr int) (int, int) {
	num := 0

	// Reads number digit by digit and stores in num.
	for {
		if digit, err := strconv.Atoi(string(input[curr])); err == nil {
			if num == 0 {
				num = digit
			} else {
				num = (num * 10) + digit
			}
			curr++
		} else {
			break
		}
	}

	return num, curr
}

// Parses and returns the following tokens from a string of the format <i-j charRule: password>:
// i int - first number in the rule.
// j int - second number in the rule.
// charRule string - char specified by rule.
// password string - a slice of input containing the password.
func Tokenize(input string) (int, int, string, string) {
	i := -1
	j := -1
	charRule:= ""
	password := ""

	// Iterate input string to parse i, j, char_rule, order password in order.
	curr := 0
	for {
		if i == -1 {
			i, curr = readNum(input, curr)
		} else if string(input[curr]) == "-" {
			curr++
		} else if j == -1 {
			j, curr = readNum(input, curr)
		} else {
			curr++
			charRule = string(input[curr])
			password = input[curr + 3:]
			break
		}
	}

	return i, j, charRule, password
}

// Solution for the Password Philosophy problem (2.1):
func PasswordPhilosophyA() (func(string), func()) {
	// For logging purposes.
	inputSize := 0
	// Counts the number of valid passwords.
	validCount:= 0

	// Processes line input and stores state in closure.
	processor := func(line string) {
		// fmt.Println(line)
		inputSize++
		min, max, charRule, password := Tokenize(line)

		// fmt.Printf("At least %d %s's and at most %d %s's in %s.\n", min, charRule, max, charRule, password)

		charCount := 0
		for _, char := range password {
			if string(char) == charRule {
				charCount++
			}
		}

		// fmt.Printf("Found %d %s's.", charCount, charRule)

		if charCount >= min && charCount <= max {
			validCount++
			// fmt.Println(" Valid.")
		} else {
			// fmt.Println(" Invalid.")
		}
	}

	// Prints current closure state.
	result := func() {
		fmt.Printf("A: There are %d/%d valid passwords.\n", validCount, inputSize)
	}

	return processor, result
}
