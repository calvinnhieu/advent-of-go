package solutions

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// Solution for the Report Repair problem (1.2):
// This 3-sum does not re-use the 2-sum implementation in 1-1 because
// 1-1 determines a solution on read. In other words, 1-1 is optimized
// to find a solution before reading the entire input.
func ReportRepairB() (func(string), func()) {
	// Int slice to store input nums
	nums := make([]int, 0)

	// Processes line input and stores state in closure.
	processor := func(line string) {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	// Finds 3-sum and prints the result.
	result := func() {
		// Find 2-sum for every (2020 - x) where (x E nums)
		// 2-sum algorithm:
		// - Maintain two indices at the start and end of the sorted slice
		// - If sum > target, increment lower
		// - If sum < target, decrement higher
		// - Repeat until target sum is found
		sort.Ints(nums)
		for i, num := range nums {
			target := 2020 - num
			j := 0
			k := len(nums) - 1
			for j < k {
				if sum := nums[j] + nums[k]; j == i || sum < target {
					j = j + 1
				} else if k == i || sum > target {
					k = k - 1
				} else {
					fmt.Printf("The sum and product of (%d %d %d) are %d and %d respectively.\n", num, nums[j], nums[k], num + nums[j] + nums[k], num * nums[j] * nums[k])
					return
				}
			}
		}
	}

	return processor, result
}
