package main

import (
	"fmt"
)

// input matrix m * n []
// 0 means for blank
// 1 means for black
// -1 means for white

// output
// 1 means black wins
// -1 means white wins
// 0 means continuing

var directions = [][][]int{{{-3, -3}, {-2, -2}, {-1, -1}, {1, 1}, {2, 2}, {3, 3}},
	{{-2, 2}, {-1, 1}, {1, -1}, {2, -2}},
	{{-2, 0}, {-1, 0}, {1, 0}, {2, 0}},
	{{0, -2}, {0, -1}, {0, 1}, {0, 2}}}

func checkMatrix(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := grid[i][j]
			if cur != 0 {
				for k := 0; k < 4; k++ {
					count := 1
					for l := 1; l < len(directions[k])-1; l++ {
						newi := i + directions[k][l][0]
						newj := j + directions[k][l][1]
						if newi < 0 || newi >= m || newj < 0 || newj >= n || grid[newi][newj] == 0 || grid[newi][newj] != cur {
							// This line doesn't meet
							break
						}
						count++
					}
					if grid[i+directions[k][0][0]][j+directions[k][0][1]] == cur || grid[i+directions[k][len(directions[k])-1][0]][j+directions[k][len(directions[k])-1][1]] == cur {
						continue
					}
					if count == 5 { // meet
						if cur == 1 {
							return 1
						} else if cur == -1 {
							return -1
						}
					}
				}
			}
		}
	}
	// Blank
	return 0
}

func main() {
	grid := [][]int{{0, 0, 0, 0, 1}, {0, 0, 0, 1, 0}, {0, -1, 1, -1, 0}, {0, 1, -1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	fmt.Println(checkMatrix(grid))
}
