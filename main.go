package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func stoneGameVII(stones []int) int {
	n := len(stones)
	var (
		record [][]int
		sum    [][]int
	)
	sum = make([][]int, n)
	record = make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, n)
		record[i] = make([]int, n)
		if i < n-1 {
			record[i][i+1] = Max(stones[i], stones[i+1])
		}
		total := stones[i]
		sum[i][i] = total
		for j := i + 1; j < n; j++ {
			total = total + stones[j]
			sum[i][j] = total
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			record[i][j] = Max(sum[i+1][j]-record[i+1][j], sum[i][j-1]-record[i][j-1])
		}
	}

	return record[0][n-1]
}
func main() {
	stones := []int{5, 3, 1, 4, 2}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(stoneGameVII(stones))
}
