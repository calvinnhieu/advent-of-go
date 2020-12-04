package solutions

import (
	"fmt"
)

// Solution for the Password Philosophy problem (2.2):
func PasswordPhilosophyB() (func(string), func()) {
	// For logging purposes.
	inputSize := 0
	// Counts the number of valid passwords.
	validCount:= 0

	// Processes line input and stores state in closure.
	processor := func(line string) {
		// fmt.Println(line)
		inputSize++
		lo, hi, charRule, password := Tokenize(line)
		i := lo - 1
		j := hi - 1

		// fmt.Printf("%s at indices %d or %d but not both in %s.", charRule, i, j, password)

		check1 := string(password[i]) == charRule 
		check2 := string(password[j]) == charRule
		if (check1 || check2) && !(check1 && check2) {
			validCount++
			// fmt.Println(" Valid.\n")
		} else {
			// fmt.Println(" Invalid.\n")
		}
	}

	// Prints current closure state.
	result := func() {
		fmt.Printf("There are %d/%d valid passwords.\n", validCount, inputSize)
	}

	return processor, result
}
