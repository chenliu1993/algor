/*
 * @lc app=leetcode.cn id=419 lang=golang
 *
 * [419] 甲板上的战舰
 */

// @lc code=start
var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func countBattleships(board [][]byte) int {
	m := len(board)
	n := len(board[0])
	var iterate func(x, y int)
	iterate = func(x, y int) {
		if board[x][y] == '.' {
			return
		}
		board[x][y] = '.'
		for i := 0; i < len(directions); i++ {
			newx := x + directions[i][0]
			newy := y + directions[i][1]
			if newx < 0 || newx > m-1 || newy < 0 || newy > n-1 {
				continue
			}
			iterate(newx, newy)
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				iterate(i, j)
				ans++
			}
		}
	}
	return ans
}

// @lc code=end

