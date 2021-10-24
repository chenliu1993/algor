/*
 * @lc app=leetcode.cn id=1905 lang=golang
 *
 * [1905] 统计子岛屿
 */

// @lc code=start
var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	var (
		m               int
		n               int
		count           int
		findChildIsland func(int, int)
	)
	m = len(grid1)
	n = len(grid1[0])
	findChildIsland = func(posiX, posiY int) {
		if grid2[posiX][posiY] == 0 {
			return
		}
		grid2[posiX][posiY] = 0
		for i := 0; i < len(directions); i++ {
			newX := posiX + directions[i][0]
			newY := posiY + directions[i][1]
			if newX >= 0 && newX < m && newY >= 0 && newY < n {
				findChildIsland(newX, newY)
			}
		}
	}
	count = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] != 0 && grid1[i][j] == 0 {
				findChildIsland(i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] != 0 && grid1[i][j] != 0 {
				findChildIsland(i, j)
				count = count + 1
			}
		}
	}
	return count
}

// @lc code=end

