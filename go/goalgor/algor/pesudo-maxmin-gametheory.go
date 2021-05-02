package algor

import "math"

const (
	CompWin  = math.MaxInt64
	Draw     = 0
	CompLoss = -1 * math.MaxInt64
)

// This program tries to find the best move for computer to win which means
// Human will be always want to minimize the value on the position
// And computer will be always want to maximize the value on the position

func FindCompMove(board [][]int, bestMove *int, value *int) {
	var (
		dc, i, resp int
	)
	if FullBoard(board) {
		*value = Draw
	} else if CompWin(board, bestMove) {
		*value = CompWin
	} else {
		*value = CompLoss
		for i := 0; i < 9; i++ {
			if isEmpty(board, i) {
				place(board, i, comp)
				FindHumanMove(board, &dc, &resp)
				unplace(board, i)

				if resp > *value {
					*value = resp
					*bestMove = i
				}
			}
		}
	}
}

func FindHumanMove(board [][]int, bestMove *int, value *int) {
	var (
		dc, i, resp int
	)
	if FullBoard(board) {
		*value = Draw
	} else if HumanWin(board, bestMove) {
		*value = CompLoss
	} else {
		*value = CompWin
		for i := 0; i < 9; i++ {
			if isEmpty(board, i) {
				place(board, i, comp)
				FindCompMove(board, &dc, &resp)
				unplace(board, i)

				if resp < *value {
					*value = resp
					*bestMove = i
				}
			}
		}
	}
}
