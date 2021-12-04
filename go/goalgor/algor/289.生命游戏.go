/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] 生命游戏
 */

// @lc code=start
var direction = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	var (
		count, newi, newj int
		ans               [][]int
	)
	ans = make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count = 0
			for k := 0; k < len(direction); k++ {
				newi = i + direction[k][0]
				newj = j + direction[k][1]
				if newi < 0 || newi > m-1 || newj < 0 || newj > n-1 {
					continue
				}
				count = count + board[newi][newj]
			}
			if board[i][j] == 1 {
				if count < 2 || count > 3 {
					ans[i][j] = 0
				} else {
					ans[i][j] = board[i][j]
				}
			} else {
				if count == 3 {
					ans[i][j] = 1
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = ans[i][j]
		}
	}
}

// @lc code=end

