/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 */

// @lc code=start
var directions = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var (
		count          int
		out            bool
		readchedBorder func(int, int)
	)
	readchedBorder = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			out = true
			return
		}
		if grid[x][y] != 1 {
			return
		}
		grid[x][y] = 0
		count++
		for i := 0; i < 4; i++ {
			readchedBorder(x+directions[i][0], y+directions[i][1])
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count = 0
				out = false
				readchedBorder(i, j)
				if !out {
					ans += count
				}
			}
		}
	}
	return ans
}

// @lc code=end

