package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"runtime"
	"github.com/calvinnhieu/advent-of-go/src/common"
)

// Approach:
// Read in input file
// Store in list, sort [O(nlogn)]
// Maintain double ended indices:
// - If < target then +i
// - Elif > target then -j
// - Else found 2-sum 
// Perform 2-sum on target = (2020 - x) for all x until 3-sum is found [O(n^2)]

func main() {
	// Int slice to store input nums
	nums := make([]int, 0)

	// Open input file
	_, src_path, _, _ := runtime.Caller(0)
	file, err := os.Open(common.Input_file(src_path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse input into nums slice and sort.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)
	
	// Handle scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Find 2-sum for every (2020 - x) where (x E nums)
	// 2-sum algorithm:
	// - Maintain two indices at the start and end of the sorted slice
	// - If sum > target, increment lower
	// - If sum < target, decrement higher
	// - Repeat until target sum is found
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
				fmt.Printf("The sum and product of (%d %d %d) are 2020 and %d respectively.\n", num, nums[j], nums[k], num * nums[j] * nums[k])
				return
			}
		}
	}
}
