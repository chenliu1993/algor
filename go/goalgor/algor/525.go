package main

import (
	"fmt"
)

func findMaxLength(nums []int) int {
	n := len(nums)
	count := 0
	maxLen := 0
	record := map[int]int{
		0: -1,
	}
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			count = count + 1
		} else {
			count = count - 1
		}
		if _, ok := record[count]; !ok {
			record[count] = i
		} else {
			if maxLen < (i - record[count]) {
				maxLen = (i - record[count])
			}
		}
	}
	fmt.Println(record)
	return maxLen
}

func main() {
	nums := []int{1, 0, 1, 0}
	fmt.Println(findMaxLength(nums))
}
