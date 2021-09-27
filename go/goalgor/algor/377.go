// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func combinationSum4(nums []int, target int) int {
// 	sort.Slice(nums, func(i, j int) bool {
// 		if nums[i] < nums[j] {
// 			return true
// 		}
// 		return false
// 	})
// 	n := len(nums)
// 	count := 0
// 	sum := 0
// 	var iterate func()
// 	iterate = func() {
// 		if sum == target {
// 			count = count + 1
// 			return
// 		}
// 		if sum > target {
// 			return
// 		}
// 		for i := 0; i < n; i++ {
// 			sum = sum + nums[i]
// 			iterate()
// 			sum = sum - nums[i]
// 		}
// 	}
// 	iterate()
// 	return count
// }

// func main() {

// 	var nums = []int{1, 2, 3}
// 	fmt.Println(combinationSum4(nums, 32))
// }

// Above over time limit
// So DP

func combinationSum4(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})
	n := len(nums)
	record := make([]int, target+1)
	record[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < n; j++ {
			if i >= nums[j] {
				record[i] = record[i] + record[i-nums[j]]
			}
		}
	}
	return record[target]
}