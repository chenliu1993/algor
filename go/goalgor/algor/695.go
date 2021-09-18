func maxAreaOfIsland(grid [][]int) int {
	dir := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	m := len(grid)
	n := len(grid[0])
	maxCount := 0
	count := 0
	var dfs func(int, int)

	dfs = func(x, y int) {
		if x < 0 || x > m-1 || y < 0 || y > n-1 {
			return
		}
		if grid[x][y] != 1 {
			return
		}
		count = count + 1
		grid[x][y] = 0
		for i := 0; i < 4; i++ {
			newx := x + dir[i][0]
			newy := y + dir[i][1]
			dfs(newx, newy)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count = 0
				dfs(i, j)
				if maxCount < count {
					maxCount = count
				}
			}
		}
	}
	return maxCount
}