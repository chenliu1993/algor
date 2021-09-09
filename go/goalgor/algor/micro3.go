package main

// m*n nodes
// n row m column
// left to right
// right step 1
// down step 1

import (
	"fmt"
)

var (
	n int
	m int
)

func main() {
	// Initialize
	n = 3
	m = 4
	// To record my steps
	walked := make([][]int, n)
	for i := 0; i < n; i++ {
		walked[i] = make([]int, m)
		walked[i][m-1] = 1
	}

	for i := 0; i < m; i++ {
		walked[n-1][i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			walked[i][j] = walked[i+1][j] + walked[i][j+1]
		}
	}

	fmt.Println(walked)
	fmt.Println(walked[0][0])
}
