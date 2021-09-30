package main

import (
	"fmt"
	"sort"
)

func canPartition(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}
	sort.Ints(nums)
	total := 0
	for i := 0; i < n; i++ {
		total = total + nums[i]
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	record := make([][]int, n)
	for i := 0; i < n; i++ {
		record[i] = make([]int, target+1)
	}
	record[0][nums[0]] = 1
	// [0....i] contains sub slice which total equals to target
	for i := 1; i < n; i++ {
		for j := 1; j < target+1; j++ {
			if j < nums[i] {
				record[i][j] = record[i-1][j]
				continue
			}
			if record[i-1][j-nums[i]] == 1 || record[i-1][j] == 1 {
				record[i][j] = 1
			}
		}
	}
	if record[n-1][target] == 1 {
		return true
	}
	return false
}

func main() {
	nums := []int{9, 5}
	fmt.Println(canPartition(nums))
}
