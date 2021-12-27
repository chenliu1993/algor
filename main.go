package main

import (
	"fmt"
)

var positions = [][]int{{-1, -2}, {1, -2}, {-1, 2}, {1, 2}, {-2, -1}, {-2, 1}, {2, -1}, {2, 1}}

func knightProbability(n int, k int, row int, column int) float64 {
	var (
		newi, newj int
		record     [][][]float64
	)

	k = k + 1
	record = make([][][]float64, k)
	for i := 0; i < k; i++ {
		record[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			record[i][j] = make([]float64, n)
		}
	}

	record[k-1][row][column] = float64(1)
	for k = k - 2; k >= 0; k-- {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for p := 0; p < len(positions); p++ {
					newi = i + positions[p][0]
					newj = j + positions[p][1]
					if newi < 0 || newi >= n || newj < 0 || newj >= n {
						continue
					}
					record[k][newi][newj] = record[k][newi][newj] + record[k+1][i][j]/float64(8)
				}
			}
		}
	}
	ans := float64(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans = ans + record[0][i][j]
		}
	}
	return ans
}

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	n := 8
	k := 3
	row := 6
	column := 4
	// 0.00019
	fmt.Println(knightProbability(n, k, row, column))

}

// var positions = [][]int{{-1, -2}, {1, -2}, {-1, 2}, {1, 2}, {-2, -1}, {-2, 1}, {2, -1}, {2, 1}}

// func knightProbability(n int, k int, row int, column int) float64 {
// 	var (
// 		counter, first, rear int
// 		record, visited      [][]int
// 		sum                  []float64
// 		q                    [][]int
// 	)
// 	record = make([][]int, n)
// 	visited = make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		record[i] = make([]int, n)
// 		visited[i] = make([]int, n)
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			for m := 0; m < len(positions); m++ {
// 				newi := i + positions[m][0]
// 				newj := j + positions[m][1]
// 				if newi < 0 || newi >= n || newj < 0 || newj >= n {
// 					continue
// 				}
// 				record[i][j]++
// 			}
// 		}
// 	}
// 	sum = make([]float64, k+1)
// 	sum[0] = float64(1)
// 	q = [][]int{{row, column}}
// 	first = 0
// 	rear = first + 1
// 	for i := 1; i <= k; i++ {
// 		counter = 0
// 		end := rear
// 		if end == first {
// 			continue
// 		}
// 		for j := first; j < end; j++ {
// 			counter = counter + record[q[j][0]][q[j][1]]
// 			for m := 0; m < len(positions); m++ {
// 				newRow := q[j][0] + positions[m][0]
// 				newColumn := q[j][1] + positions[m][1]
// 				if newRow < 0 || newRow >= n || newColumn < 0 || newColumn >= n {
// 					continue
// 				}
// 				if visited[newRow][newColumn] != first+1 {
// 					q = append(q, []int{newRow, newColumn})
// 					visited[newRow][newColumn] = first + 1
// 					rear++
// 				}
// 			}
// 		}
// 		sum[i] = sum[i-1] * float64(counter) / float64(8*(end-first))
// 		first = end
// 	}
// 	return sum[k]
// }
