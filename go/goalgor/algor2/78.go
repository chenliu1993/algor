package main

import (
	"sort"
)

func basicSubsets(nums []int) {
	for len(nums) == 0 {
		ans = append(ans, targets)
		return
	}
	old_targets := make([]int, len(targets))
	copy(old_targets, targets)
	targets = append(targets, nums[0])
	basicSubsets(nums[1:])
	targets = make([]int, len(old_targets))
	copy(targets, old_targets)

	basicSubsets(nums[1:])
}

var (
	targets []int
	ans     [][]int
)

func subsets(nums []int) [][]int {
	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] < nums[j]
	// })
	sort.Ints(nums)
	targets = []int{}
	ans = [][]int{}
	basicSubsets(nums)
	return ans
}

// func main() {
// 	nums := []int{1, 2, 3}
// 	fmt.Println(subsets(nums))

// }
