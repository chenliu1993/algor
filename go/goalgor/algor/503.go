package main

import (
	"fmt"
)

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := []int{}
	// Iterate through every element
	for i := 0; i < n; i++ {
		target := nums[i]
		j := i + 1
		j = j % n
		for ; j != i; j = j % n {
			if nums[j] > target {
				ans = append(ans, nums[j])
				break
			} else {
				j = j + 1
			}
		}
		if j == i {
			ans = append(ans, -1)
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums))
}
