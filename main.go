package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func findMaxLength(nums []int) int {
	n := len(nums)
	var (
		sum    int
		maxLen int
		dict   map[int]int
		record map[int]int
	)
	maxLen = 0
	record = map[int]int{
		0: -1,
	}
	dict = map[int]int{
		1: 1,
		0: -1,
	}
	sum = 0
	for i := 0; i < n; i++ {
		sum = sum + dict[nums[i]]
		if _, ok := record[sum]; !ok {
			record[sum] = i
		} else {
			maxLen = Max(maxLen, i-record[sum])
		}
	}
	return maxLen
}

func main() {
	nums := []int{0, 1}
	// nums := []int{0, 1}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(findMaxLength(nums))
}

// Over time limit
// func findMaxLength(nums []int) int {
// 	n := len(nums)
// 	var (
// 		maxLen int
// 		sums   []int
// 	)
// 	maxLen = 0
// 	sums = make([]int, n)
// 	sums[0] = nums[0]
// 	for i := 1; i < n; i++ {
// 		sums[i] = sums[i-1] + nums[i]
// 		if sums[i]*2 == i+1 {
// 			maxLen = Max(maxLen, i+1)
// 			continue
// 		}
// 		for j := i - 1; j >= 0; j-- {
// 			if (sums[i]-sums[j])*2 == i-j {
// 				maxLen = Max(maxLen, i-j)
// 			}
// 		}
// 	}
// 	return maxLen
// }
