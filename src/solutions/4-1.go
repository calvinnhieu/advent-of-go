package solutions

import (
	"fmt"
	// "strconv"
)

// Parses n keys from <key:field> pairs in a space-separated string.
// Returns a slice of n key fields.
// Assumes key lengths are fixed at 3.
func getKeys(line string) []string {
	keys := make([]string, 0)

	// Searches for ":" to determine key location and backtracks.
	// Appends to keys slice.
	for i, curr := range line {
		if string(curr) == ":" {
			keys = append(keys, line[i - 3:i])
		}
	}

	// fmt.Println(keys)

	return keys
}

// Reads a slice of keys and sets the corresponding bit for each key.
// Returns the modified bit tracker.
func setKeys(keys []string, bitTracker uint8) uint8 {
	for _, key := range keys {
		switch key {
		case "byr":
			bitTracker = bitTracker + (1 << 0)
		case "iyr":
			bitTracker = bitTracker + (1 << 1)
		case "eyr":
			bitTracker = bitTracker + (1 << 2)
		case "hgt":
			bitTracker = bitTracker + (1 << 3)
		case "hcl":
			bitTracker = bitTracker + (1 << 4)
		case "ecl":
			bitTracker = bitTracker + (1 << 5)
		case "pid":
			bitTracker = bitTracker + (1 << 6)
		case "default":
			continue
		}
		// fmt.Println(strconv.FormatInt(int64(bitTracker), 2))
	}
	return bitTracker
}

// Solution for the PassportProcessing problem (4.1):
func PassportProcessingA() (func(string), func()) {
	// Tracks presence of 8 keys of the current passport entry.
	var keyTracker uint8 = 0
	// Counts the number of valid passports.
	var validCount int = 0

	// Processes line input and stores state in closure.
	processor := func(line string) {
		// If non-empty line:
		// - Parse fields and track in 8 bits. Ignore cid
		if line != "" {
			keyTracker = setKeys(getKeys(line), keyTracker)
		// If empty line:
		// - Check if all 8 or 7 (-cid) bits are set. Count valids.
		// - Reset bits
		} else {
			// fmt.Print(keyTracker)
			if keyTracker == 127 {
				validCount++
				// fmt.Println("VALID")
			} else {
				// fmt.Println("INVALID")
			}
			keyTracker = 0
			// fmt.Println("")
		}
	}

	// Prints current closure state.
	result := func() {
		fmt.Printf("A: There are %d \"valid\" passports.\n", validCount)
	}

	return processor, result
}
