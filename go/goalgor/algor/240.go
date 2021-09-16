var direction = [][]int{{1, 0}, {0, 1}}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	count := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		visited[x][y] = 1
		if matrix[x][y] == target {
			count = count + 1
			return
		}
		for i := 0; i < 2; i++ {
			newx := x + direction[i][0]
			newy := y + direction[i][1]
			if newx > m-1 || newy > n-1 || visited[newx][newy] == 1 || matrix[newx][newy] > target {
				continue
			}
			dfs(newx, newy)
		}
	}
	dfs(0, 0)
	if count == 0 {
		return false
	}
	return true
}
