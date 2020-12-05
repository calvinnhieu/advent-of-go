package solutions

import (
	"fmt"
	"strconv"
)

// A set containing all valid eye color values.
// Implemented as a map of empty structs.
// Note: empty structs do not require additional memory.
var ecls = map[string]struct{}{"amb": {}, "blu": {}, "brn": {}, "gry": {}, "grn":{}, "hzl": {}, "oth":{}}

// Determines whether a string contains a number that is within the supplied range.
func valueInRange(value string, min, max int) bool {
	num, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return num >= min && num <= max
}

// Determines whether a string contains a valid height defined by:
// Between 150cm and 193cm or 59in and 76in (inclusive).
func isValidHeight(value string) bool {
	hgtMeasurement := value[:len(value) - 2]
	hgtUnit := value[len(value) - 2:len(value)]

	return (hgtUnit == "cm" && valueInRange(hgtMeasurement, 150, 193)) || (hgtUnit == "in" && valueInRange(hgtMeasurement, 59, 76))
}

// Determines whether a string contains a valid hex value
// beginning with a "#" and followed by (6) 0-9, a-f.
func isHex(value string) bool {
	if string(value[0]) != "#" {
		return false
	}

	var char string
	for _, r := range value[1:] {
		char = string(r)
		if !((char >= "a" && char <= "f") || (char >= "0" && char <= "9")) {
			return false
		}
	}

	return true
}

// Parses n KVPs from <key:val> pairs in a space-separated string.
// Returns a map of n KVPs.
// Assumes constant key lengths (3) and constant value lengths per field.
func getKvps(line string) map[string]string {
	i := 0
	kvps := make(map[string]string)

	// Iterates string and searches for ":" to identify KVP.
	// Backtracks fixed length 3 for key, iterates to next space for value.
	var curr string
	var valueStart int
	for i < len(line) {
		curr = string(line[i])
		// Found key.
		if curr == ":" {
			key := line[i - 3:i]
			i++
			valueStart = i
			// Iterate to retrieve value.
			for i < len(line) {
				curr = string(line[i])
				if curr == " " {
					break
				}
				i++
			}

			kvps[key] = line[valueStart:i]
		} else {
			i++
		}
	}

	// fmt.Println(kvps)

	return kvps
}

// Reads a map of KVPs sets the corresponding bit if valid.
// Returns the modified bit tracker.
func validateKvps(kvps map[string]string, bitTracker uint8) uint8 {
	for key, val := range kvps {
		switch key {
		case "byr":
			if valueInRange(val, 1920, 2002) {
				bitTracker = bitTracker + (1 << 0)
			}
		case "iyr":
			if valueInRange(val, 2010, 2020) {
				bitTracker = bitTracker + (1 << 1)
			}
		case "eyr":
			if valueInRange(val, 2020, 2030) {
				bitTracker = bitTracker + (1 << 2)
			}
		case "hgt":
			if isValidHeight(val) {
				bitTracker = bitTracker + (1 << 3)
			}
		case "hcl":
			if isHex(val) {
				bitTracker = bitTracker + (1 << 4)
			}
		case "ecl":
			if _, ok := ecls[val]; ok {
				bitTracker = bitTracker + (1 << 5)
			}
		case "pid":
			if _, err := strconv.Atoi(val); err == nil && len(val) == 9 {
				bitTracker = bitTracker + (1 << 6)
			}
		default:
			continue
		}
		//fmt.Println(strconv.FormatInt(int64(bitTracker), 2))
	}
	return bitTracker
}

// Solution for the PassportProcessing problem (4.2):
func PassportProcessingB() (func(string), func()) {
	// Tracks validity of 7 KVPs of the current passport entry.
	var kvpTracker uint8 = 0
	// Counts the number of valid passports.
	var validCount int = 0

	// Processes line input and stores state in closure.
	processor := func(line string) {
		// If non-empty line:
		// - Validate fields and track in 7 bits. Ignore cid.
		if line != "" {
			kvpTracker = validateKvps(getKvps(line), kvpTracker)
		// If empty line:
		// - Check if all 7 bits are set. Count valids.
		// - Reset bits
		} else {
			// fmt.Print(keyTracker)
			if kvpTracker == 127 {
				validCount++
				// fmt.Println("VALID")
			} else {
				// fmt.Println("INVALID")
			}
			kvpTracker = 0
			// fmt.Println("")
		}
	}

	// Prints current closure state.
	result := func() {
		fmt.Printf("A: There are %d \"valid\" passports.\n", validCount)
	}

	return processor, result
}

