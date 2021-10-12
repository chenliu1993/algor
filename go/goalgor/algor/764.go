package main

import "fmt"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	m := len(mines)
	// Since initialized matrix is filled by 0, then here I will find 0
	matrix := make([][]int, n)
	record1 := make([][]int, n)
	record2 := make([][]int, n)
	record3 := make([][]int, n)
	record4 := make([][]int, n)
	for i := 0; i < n; i++ {
		record1[i] = make([]int, n)
		record2[i] = make([]int, n)
		record3[i] = make([]int, n)
		record4[i] = make([]int, n)
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		matrix[mines[i][0]][mines[i][1]] = 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == 0 {
					record1[i][j] = 1
					record2[i][j] = 1
				} else {
					record1[i][j] = 0
					record2[i][j] = 0
				}
			} else if matrix[i][j] == 1 {
				record1[i][j] = 0
				record2[i][j] = 0
			} else {
				record1[i][j] = record1[i-1][j] + 1
				record2[i][j] = record2[i][j-1] + 1
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == n-1 || j == n-1 {
				if matrix[i][j] == 0 {
					record3[i][j] = 1
					record4[i][j] = 1
				} else {
					record3[i][j] = 0
					record4[i][j] = 0
				}
			} else if matrix[i][j] == 1 {
				record3[i][j] = 0
				record4[i][j] = 0
			} else {
				record3[i][j] = record3[i+1][j] + 1
				record4[i][j] = record4[i][j+1] + 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			record := Min(record1[i][j], record2[i][j])
			record = Min(record, record3[i][j])
			record = Min(record, record4[i][j])
			ans = Max(ans, record)
		}
	}
	// fmt.Println(record1, record2)
	return ans
}

func main() {
	// mines := [][]int{{3, 0}, {3, 3}}
	n := 2
	mines := [][]int{{0, 1}, {1, 0}, {1, 1}}
	fmt.Println(orderOfLargestPlusSign(n, mines))
}
