package main

import (
	"fmt"
)

func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
	stk := []int{}
	ans := []int{}

	removeCount := n - k
	for i := 0; i < n; i++ {
		for removeCount != 0 && len(stk) != 0 && nums[stk[len(stk)-1]] > nums[i] {
			stk = stk[:len(stk)-1]
			removeCount = removeCount - 1
		}
		stk = append(stk, i)
	}
	for i := 0; i < removeCount; i++ {
		stk = stk[:len(stk)-1]
	}
	for i := 0; i < k; i++ {
		ans = append(ans, nums[stk[i]])
	}

	return ans
}

func main() {
	nums := []int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2}
	k := 3
	// nums := []int{3, 5, 2, 6}
	// k := 2
	fmt.Println(mostCompetitive(nums, k))
}
