/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
var directions = [][]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 0}, {-1, -1}, {-1, 1}}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	m := len(board)
	n := len(board[0])
	var (
		first, rear    int
		queue, visited [][]int
	)

	queue = [][]int{{click[0], click[1]}}
	visited = make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	first = 0
	rear = first + 1
	for first != rear {
		end := rear
		for i := first; i < end; i++ {
			r := queue[i][0]
			c := queue[i][1]
			if board[r][c] != 'E' {
				continue
			}
			mines := 0
			for i := 0; i < len(directions); i++ {
				newr := r + directions[i][0]
				newc := c + directions[i][1]
				if newr < 0 || newr >= m || newc < 0 || newc >= n {
					continue
				}
				if board[newr][newc] == 'M' {
					mines++
				}
			}
			if mines != 0 {
				board[r][c] = byte('0' + mines)
			} else {
				board[r][c] = 'B'
				for i := 0; i < len(directions); i++ {
					newr := r + directions[i][0]
					newc := c + directions[i][1]
					if newr < 0 || newr >= m || newc < 0 || newc >= n {
						continue
					}
					if visited[newr][newc] == 0 {
						queue = append(queue, []int{newr, newc})
						visited[newr][newc] = 1
						rear++
					}
				}
			}
		}
		first = end
	}
	return board
}

// @lc code=end

