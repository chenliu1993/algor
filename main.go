package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	n := len(nums)
	sums := make([]int, n)
	sums[0] = nums[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	count := 0
	for i := 0; i < n; i++ {
		if sums[i] == k {
			count++
		}
		for j := i - 1; j >= 0; j-- {
			if sums[i]-sums[j] == k {
				count++
			}
		}
	}
	return count
}

func main() {
	k := 3
	nums := []int{1, 2, 3}
	// nums := []int{1, 2, 3}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(subarraySum(nums, k))
}
