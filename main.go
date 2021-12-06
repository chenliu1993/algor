package main

import (
	"fmt"
)

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	var (
		maxLen        int
		length        int
		first, second int
		res           int
		exists        map[int]bool
	)
	exists = map[int]bool{}
	for i := 0; i < n; i++ {
		exists[arr[i]] = true
	}
	maxLen = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			length = 2
			first = arr[i]
			second = arr[j]
			res = first + second
			for exists[res] {
				length++
				first = second
				second = res
				res = first + second
			}
			maxLen = Max(maxLen, length)
		}
	}
	if maxLen == 2 {
		return 0
	}
	return maxLen
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// arr := []int{1, 3, 7, 11, 12, 14, 18}
	// arr := []int{2, 4, 5, 6, 7, 8, 11, 13, 14, 15, 21, 22, 34}
	fmt.Println(lenLongestFibSubseq(arr))
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
}
