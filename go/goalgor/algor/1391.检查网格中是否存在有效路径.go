/*
 * @lc app=leetcode.cn id=1391 lang=golang
 *
 * [1391] 检查网格中是否存在有效路径
 */

// @lc code=start

func getRoute(x int) [][]int {
	if x == 1 {
		return [][]int{{0, 1, 1}, {0, 1, 3}, {0, 1, 5}, {0, -1, 1}, {0, -1, 4}, {0, -1, 6}}
	} else if x == 2 {
		return [][]int{{1, 0, 2}, {1, 0, 5}, {1, 0, 6}, {-1, 0, 2}, {-1, 0, 3}, {-1, 0, 4}}
	} else if x == 3 {
		return [][]int{{0, -1, 1}, {0, -1, 4}, {0, -1, 6}, {1, 0, 2}, {1, 0, 5}, {1, 0, 6}}
	} else if x == 4 {
		return [][]int{{0, 1, 1}, {0, 1, 3}, {0, 1, 5}, {1, 0, 2}, {1, 0, 5}, {1, 0, 6}}
	} else if x == 5 {
		return [][]int{{0, -1, 1}, {0, -1, 4}, {0, -1, 6}, {-1, 0, 2}, {-1, 0, 3}, {-1, 0, 4}}
	}
	return [][]int{{0, 1, 1}, {0, 1, 3}, {0, 1, 5}, {-1, 0, 2}, {-1, 0, 3}, {-1, 0, 4}}
}

func hasValidPath(grid [][]int) bool {
	var (
		m       int
		n       int
		flag    bool
		iterate func(int, int)
	)
	m = len(grid)
	n = len(grid[0])
	if m == 1 && n == 1 {
		return true
	}
	iterate = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == 0 {
			return
		}
		route := getRoute(grid[x][y])
		temp := grid[x][y]
		grid[x][y] = 0
		for i := 0; i < len(route); i++ {
			newx := x + route[i][0]
			newy := y + route[i][1]
			if newx >= 0 && newx < m && newy >= 0 && newy < n && route[i][2] != grid[newx][newy] {
				continue
			}
			if newx == m-1 && newy == n-1 {
				flag = true
				return
			}
			iterate(newx, newy)
		}
		grid[x][y] = temp
	}
	flag = false
	iterate(0, 0)
	return flag
}

// @lc code=end

