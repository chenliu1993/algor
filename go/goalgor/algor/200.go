var direction = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	var dfs func(int, int, [][]byte, [][]int)
	dfs = func(x, y int, grid [][]byte, visited [][]int) {
		grid[x][y] = '0'
		visited[x][y] = 1
		for i := 0; i < len(direction); i++ {
			newx := x + direction[i][0]
			newy := y + direction[i][1]
			if newx < 0 || newx >= m || newy < 0 || newy >= n || visited[newx][newy] == 1 {
				continue
			}
			if grid[newx][newy] == '1' {
				dfs(newx, newy, grid, visited)
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j, grid, visited)
				count = count + 1
			}
		}
	}
	return count
}
