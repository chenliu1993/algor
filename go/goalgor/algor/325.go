package main

import (
	"fmt"
)

// Just a possible way, similar to 525, never tested on leetcode since 325 is a plus problem
func findMaxLength(nums []int, k int) int {
	n := len(nums)
	count := 0
	maxLen := 0
	record := map[int]int{
		0: -1,
	}
	for i := 0; i < n; i++ {
		count = count + nums[i]
		if _, ok := record[count]; !ok {
			record[count] = i
		}
		if _, ok := record[count-k]; ok {
			if maxLen < i-record[count-k] {
				fmt.Println(i, count, record[count-k])
				maxLen = i - record[count-k]
			}
		}
	}
	return maxLen
}

func main() {
	nums := []int{1, 0, 1, 0, 1}
	k := 3
	fmt.Println(findMaxLength(nums, k))
}
