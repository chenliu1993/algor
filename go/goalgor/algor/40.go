package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i] < candidates[j] {
			return true
		}
		return false
	})
	dict := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if _, ok := dict[candidates[i]]; !ok {
			dict[candidates[i]] = 1
		} else {
			dict[candidates[i]] = dict[candidates[i]] + 1
		}
	}

	newCandidates := []int{}
	for k := range dict {
		if k <= target {
			newCandidates = append(newCandidates, k)
		}
	}
	n = len(newCandidates)

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
			if dict[newCandidates[i]] <= 0 {
				continue
			}
			sum = sum + newCandidates[i]
			dict[newCandidates[i]] = dict[newCandidates[i]] - 1
			result = append(result, newCandidates[i])
			iterate(i)
			sum = sum - newCandidates[i]
			dict[newCandidates[i]] = dict[newCandidates[i]] + 1
			result = result[:len(result)-1]
		}
	}
	iterate(0)
	return results
}

func main() {
	can := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(can, target))
}
