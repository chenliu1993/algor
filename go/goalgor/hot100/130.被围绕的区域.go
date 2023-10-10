/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start

func solve(board [][]byte) {
	var direction [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := [][]int{}

	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
			board[i][0] = '1'
		}
		if board[i][len(board[0])-1] == 'O' {
			queue = append(queue, []int{i, len(board[0]) - 1})
			board[i][len(board[0])-1] = '1'
		}
	}

	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			queue = append(queue, []int{0, i})
			board[0][i] = '1'
		}
		if board[len(board)-1][i] == 'O' {
			queue = append(queue, []int{len(board) - 1, i})
			board[len(board)-1][i] = '1'
		}
	}

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			x := start[0] + direction[k][0]
			y := start[1] + direction[k][1]
			if x < 0 || y < 0 || x > len(board)-1 || y > len(board[0])-1 || board[x][y] != 'O' {
				continue
			}
			queue = append(queue, []int{x, y})
			board[x][y] = '1'
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '1' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}

}

// @lc code=end

