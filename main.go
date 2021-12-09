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
		maxLen int
		dict   map[int]int
		record [][]int
	)
	dict = map[int]int{}
	record = make([][]int, n)
	maxLen = 0
	for i := 0; i < n; i++ {
		dict[arr[i]] = i
		record[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			k, ok := dict[arr[i]-arr[j]]
			if (arr[i]-arr[j] < arr[j]) && ok {
				record[j][i] = record[k][j] + 1
				maxLen = Max(maxLen, record[j][i]+2)
			}
		}
	}
	if maxLen < 3 {
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
