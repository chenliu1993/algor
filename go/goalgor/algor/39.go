package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i] < candidates[j] {
			return true
		}
		return false
	})
	results := [][]int{}
	result := []int{}
	sum := 0
	var iterate func(int)
	iterate = func(start int) {
		if sum > target {
			return
		}
		if sum == target {
			tempResult := make([]int, len(result))
			copy(tempResult, result)
			results = append(results, tempResult)
			return
		}
		for i := start; i < n; i++ {
			if candidates[i] > target {
				continue
			}
			sum = sum + candidates[i]
			result = append(result, candidates[i])
			iterate(i)
			sum = sum - candidates[i]
			result = result[:len(result)-1]
		}
	}
	iterate(0)
	return results
}

func main() {
	can := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(can, target))
}
