/*
 * @lc app=leetcode.cn id=1319 lang=golang
 *
 * [1319] 连通网络的操作次数
 */

// @lc code=start
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func makeConnected(n int, connections [][]int) int {
	m := len(connections)
	if m < n-1 {
		return -1
	}
	var (
		clusters   int
		maxDegrees int
		visited    []int
		adjs       [][]int
		dfs        func(int)
	)

	dfs = func(x int) {
		visited[x] = 1
		for i := 0; i < len(adjs[x]); i++ {
			if visited[adjs[x][i]] == 0 {
				dfs(adjs[x][i])
			}
		}
	}
	adjs = make([][]int, n)
	for i := 0; i < n; i++ {
		adjs[i] = []int{}
	}
	maxDegrees = 0
	for i := 0; i < m; i++ {
		adjs[connections[i][0]] = append(adjs[connections[i][0]], connections[i][1])
		adjs[connections[i][1]] = append(adjs[connections[i][1]], connections[i][0])
		maxDegrees = Max(maxDegrees, Max(len(adjs[connections[i][0]]), len(adjs[connections[i][1]])))
	}
	clusters = 0
	visited = make([]int, n)
	for i := 0; i < n; i++ {
		if visited[i] != 1 {
			visited[i] = 1
			dfs(i)
			clusters++
		}
	}

	return clusters - 1
}

// @lc code=end

