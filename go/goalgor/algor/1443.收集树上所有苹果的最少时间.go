/*
 * @lc app=leetcode.cn id=1443 lang=golang
 *
 * [1443] 收集树上所有苹果的最少时间
 */

// @lc code=start
func minTime(n int, edges [][]int, hasApple []bool) int {
	var (
		adjs    [][]int
		visited []int
		times   func(int) int
	)
	// Construct adjacent table
	adjs = make([][]int, n)
	for i := 0; i < n; i++ {
		adjs[i] = []int{}
	}
	visited = make([]int, n)
	visited[0] = 1
	for i := 0; i < len(edges); i++ {
		adjs[edges[i][0]] = append(adjs[edges[i][0]], edges[i][1])
		adjs[edges[i][1]] = append(adjs[edges[i][1]], edges[i][0])
	}
	times = func(x int) int {
		if len(adjs[x]) == 0 {
			return 0
		}
		sum := 0
		for i := 0; i < len(adjs[x]); i++ {
			if visited[adjs[x][i]] == 1 {
				continue
			}
			visited[adjs[x][i]] = 1
			time := times(adjs[x][i])
			if hasApple[adjs[x][i]] || time != 0 {
				sum = sum + 2 + time
			}
		}
		return sum
	}
	if n == 1 {
		return 0
	}
	return times(0)
}

// @lc code=end

// func minTime(n int, edges [][]int, hasApple []bool) int {
// 	var (
// 		adjs  [][]int
// 		times func(int) int
// 	)
// 	// Construct adjacent table
// 	adjs = make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		adjs[i] = make([]int, n)
// 	}
// 	for i := 0; i < len(edges); i++ {
// 		adjs[edges[i][0]][edges[i][1]] = 1
// 		adjs[edges[i][1]][edges[i][0]] = 1
// 	}
// 	times = func(x int) int {
// 		sum := 0
// 		for i := 0; i < n; i++ {
// 			if adjs[x][i] == 1 {
// 				adjs[x][i] = 0
// 				adjs[i][x] = 0
// 				time := times(i)
// 				if hasApple[i] || time != 0 {
// 					sum = sum + 2 + time
// 				}
// 			}
// 		}
// 		return sum
// 	}
// 	if n == 1 {
// 		return 0
// 	}
// 	return times(0)
// }

