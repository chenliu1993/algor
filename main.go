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

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var (
		ans         int
		west, south []int
	)
	west, south = make([]int, m), make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			south[i] = Max(south[i], grid[j][i])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			west[i] = Max(west[i], grid[i][j])
		}
	}

	fmt.Println(west, south)
	ans = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := Min(west[i], south[j])
			ans += cur - grid[i][j]
		}
	}
	return ans
}

func main() {
	// trips := [][]int{{8, 2, 3}, {4, 1, 3}, {1, 3, 6}, {8, 4, 6}, {4, 4, 8}}
	// grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	grid := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}
