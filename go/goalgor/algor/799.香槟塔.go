/*
 * @lc app=leetcode.cn id=799 lang=golang
 *
 * [799] 香槟塔
 */

// @lc code=start
func champagneTower(poured int, query_row int, query_glass int) float64 {
	var (
		more, record [][]float64
	)
	record = make([][]float64, query_row+1)
	more = make([][]float64, query_row+1)
	for i := 0; i < query_row+1; i++ {
		record[i] = make([]float64, query_glass+1)
		more[i] = make([]float64, query_glass+1)
	}

	if poured <= 1 {
		record[0][0] = float64(poured)
		more[0][0] = 0.0
	} else {
		record[0][0] = 1.0
		more[0][0] = float64(poured - 1)
	}

	for i := 1; i < query_row+1; i++ {
		for j := 0; j < query_glass+1; j++ {
			if j == 0 {
				record[i][j] = more[i-1][j] / 2
			} else if j == i {
				record[i][j] = more[i-1][j-1] / 2
			} else {
				record[i][j] = (more[i-1][j-1] + more[i-1][j]) / 2
			}
			if record[i][j] <= 1 {
				more[i][j] = float64(0)
			} else {
				more[i][j] = record[i][j] - 1
				record[i][j] = float64(1)
			}
		}
	}
	return record[query_row][query_glass]
}

// @lc code=end

