/*
 * @lc app=leetcode.cn id=1139 lang=golang
 *
 * [1139] 最大的以 1 为边界的正方形
 */

// @lc code=start
func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
func largest1BorderedSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var (
		curLen int
		maxLen int
		right  [][]int
		down   [][]int
	)
	right = make([][]int, m)
	down = make([][]int, m)
	for i := 0; i < m; i++ {
		right[i] = make([]int, n)
		down[i] = make([]int, n)
	}
	maxLen = 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 0 {
				continue
			}
			if j == n-1 {
				right[i][j] = 1
			} else {
				right[i][j] = 1 + right[i][j+1]
			}
			if i == m-1 {
				down[i][j] = 1
			} else {
				down[i][j] = 1 + down[i+1][j]
			}
			curLen = Min(right[i][j], down[i][j])
			if curLen < maxLen {
				continue
			}
			for l := curLen; l > maxLen; l-- {
				if i+l-1 < m && right[i+l-1][j] >= l && j+l-1 < n && down[i][j+l-1] >= l {
					maxLen = Max(maxLen, l)
				}
			}
		}
	}
	return maxLen * maxLen
}

// @lc code=end

