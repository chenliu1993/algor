/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	var (
		n         int
		provinces int
		dfs       func(int)
	)
	dfs = func(x int) {
		for i := 0; i < n; i++ {
			if isConnected[x][i] == 1 {
				isConnected[i][x] = 0
				isConnected[x][i] = 0
				isConnected[i][i] = 0
				if x != i {
					dfs(i)
				}
			}
		}
	}
	n = len(isConnected)
	provinces = 0
	for i := 0; i < n; i++ {
		if isConnected[i][i] == 1 {
			isConnected[i][i] = 0
			dfs(i)
			provinces++
		}
	}
	return provinces
}

// @lc code=end

