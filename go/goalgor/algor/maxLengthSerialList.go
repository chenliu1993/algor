package main

import (
	"fmt"
)

// input: []int
// longest 子序列 []int

// [2,6,8,3,4,5,1]
func maxLength(input []int) int {
	length := len(input)
	if length < 2 {
		return length
	}

	//To records subslice i:j which ends with j
	records := make([][]int, length)

	for i := 0; i < length; i++ {
		records[i] = make([]int, length)
		// Self longest is 1
		for j := i; j < length; j++ {
			records[i][j] = 1
		}
	}
	var curMin int
	maxLen := 1
	for j := 0; j < length; j++ {
		curMin = input[j]
		for i := j; i >= 1; i-- {
			if input[i-1] < curMin {
				curMin = input[i-1]
				records[i-1][j] = records[i][j] + 1
			} else {
				records[i-1][j] = records[i][j]
			}
		}
		if maxLen < records[0][j] {
			maxLen = records[0][j]
		}
	}
	return maxLen
}

func main() {
	var input = []int{2, 6, 8, 3, 4, 5, 1}
	fmt.Println(input)
	fmt.Println(maxLength(input))
}
