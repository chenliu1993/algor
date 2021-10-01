// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func countZeroAndOne(str string) []int {
// 	return []int{strings.Count(str, "0"), strings.Count(str, "1")}
// }

// func findMaxForm(strs []string, m int, n int) int {
// 	strsLen := len(strs)
// 	var (
// 		zeroSum int
// 		oneSum  int
// 		maxLen  int
// 		result  = []string{}
// 	)
// 	var iterate func(int)
// 	iterate = func(start int) {
// 		if zeroSum > m || oneSum > n || start > strsLen {
// 			return
// 		}
// 		if zeroSum <= m && oneSum <= n {
// 			if maxLen < len(result) {
// 				maxLen = len(result)
// 			}
// 			if start == strsLen {
// 				return
// 			}
// 		}

// 		cur := countZeroAndOne(strs[start])
// 		zeroSum += cur[0]
// 		oneSum += cur[1]
// 		result = append(result, strs[start])
// 		iterate(start + 1)
// 		zeroSum -= cur[0]
// 		oneSum -= cur[1]
// 		result = result[:len(result)-1]
// 		iterate(start + 1)
// 	}
// 	iterate(0)
// 	return maxLen
// }

// func main() {
// 	nums := []string{"0", "11", "1000", "01", "0", "101", "1", "1", "1", "0", "0", "0", "0", "1", "0", "0110101", "0", "11", "01", "00", "01111", "0011", "1", "1000", "0", "11101", "1", "0", "10", "0111"}
// 	fmt.Println(findMaxForm(nums, 9, 80))
// }

// Above over time limit

func countZeroAndOne(str string) []int {
	return []int{strings.Count(str, "0"), strings.Count(str, "1")}
}

func findMaxForm(strs []string, m int, n int) int {
	strsLen := len(strs)
	record := make([][][]int, strsLen)
	for i := 0; i < strsLen; i++ {
		record[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			record[i][j] = make([]int, n+1)
		}
	}
	first := countZeroAndOne(strs[0])
	if first[0] <= m && first[1] <= n {
		for i := first[0]; i <= m; i++ {
			for j := first[1]; j <= n; j++ {
				record[0][i][j] = 1
			}
		}
	}
	for k := 1; k < strsLen; k++ {
		cur := countZeroAndOne(strs[k])
		for i := 0; i < m+1; i++ {
			for j := 0; j < n+1; j++ {
				if i < cur[0] || j < cur[1] {
					record[k][i][j] = record[k-1][i][j]
				} else {
					if record[k-1][i][j] < record[k-1][i-cur[0]][j-cur[1]]+1 {
						record[k][i][j] = record[k-1][i-cur[0]][j-cur[1]] + 1
					} else {
						record[k][i][j] = record[k-1][i][j]
					}
				}
			}
		}
	}
	return record[strsLen-1][m][n]
}