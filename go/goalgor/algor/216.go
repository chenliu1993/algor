package main

import (
	"fmt"
)

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func combinationSum3(k int, n int) [][]int {

	length := len(nums)
	results := [][]int{}
	result := []int{}
	sum := 0
	var iterate func(int)
	iterate = func(start int) {
		if sum == n && len(result) == k {
			fmt.Println(sum, result)
			tempResult := make([]int, len(result))
			copy(tempResult, result)
			results = append(results, tempResult)
		}
		for i := start; i < length; i++ {
			sum = sum + nums[i]
			result = append(result, nums[i])

			iterate(i + 1)
			result = result[:len(result)-1]
			sum = sum - nums[i]
		}
	}
	iterate(0)
	return results
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
