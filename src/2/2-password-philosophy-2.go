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
// - Check char presence in specified indices (exclusive or).
// - Count and yield number of valid passwords.

// Reads a number token from a string and advances the given index.
// Returns the parsed number and new index.
// Assumes the starting index is a number, else returns 0.
// Assumes the read number does not reach the end of input, else fails.
func read_num(input string, curr int) (int, int) {
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

// Parses and returns the following tokens from a password rule string:
// min int - minimum occurrence boundary.
// max int - maximum occurrence boundary.
// char_rule string - char specified by rule.
// password string - a slice of input containing the password.
func tokenize(input string) (int, int, string, string) {
	i := -1
	j := -1
	char_rule := ""
	password := ""

	// Iterate input string to parse i, j, char_rule, order password in order.
	curr := 0
	for {
		if i == -1 {
			i, curr = read_num(input, curr)
		} else if string(input[curr]) == "-" {
			curr++
		} else if j == -1 {
			j, curr = read_num(input, curr)
		} else {
			curr++
			char_rule = string(input[curr])
			password = input[curr + 3:]
			break
		}
	}
	
	return i - 1, j - 1, char_rule, password
}

func main() {
	// For logging purposes.
	input_size := 0
	// Stores number of valid passwords.
	valid_count := 0

	// Open input file.
	_, src_path, _, _ := runtime.Caller(0)
	file, err := os.Open(common.Input_file(src_path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read input line by line.
	// Tokenize each line.
	// Determine validity by checking indices i and j for specified char (exclusive or).
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input_size++
		line := string(scanner.Text())

		// fmt.Println(line)

		i, j, char_rule, pwd := tokenize(line)

		// fmt.Printf("%s at indices %d or %d but not both in %s.", char_rule, i, j, pwd)
		
		check1 := string(pwd[i]) == char_rule
		check2 := string(pwd[j]) == char_rule
		if (check1 || check2) && !(check1 && check2) {
			valid_count++
			// fmt.Println(" Valid.\n")
		} else {
			// fmt.Println(" Invalid.\n")
		}
	}

	fmt.Printf("There are %d/%d valid passwords.\n", valid_count, input_size)
}
