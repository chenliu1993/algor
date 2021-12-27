/*
 * @lc app=leetcode.cn id=688 lang=golang
 *
 * [688] “马”在棋盘上的概率
 */

// @lc code=start
var positions = [][]int{{-1, -2}, {1, -2}, {-1, 2}, {1, 2}, {-2, -1}, {-2, 1}, {2, -1}, {2, 1}}

func knightProbability(n int, k int, row int, column int) float64 {
	var (
		ans               float64
		record, recordOld [][]float64
	)
	record = make([][]float64, n)
	recordOld = make([][]float64, n)
	for i := 0; i < n; i++ {
		record[i] = make([]float64, n)
		recordOld[i] = make([]float64, n)
	}
	recordOld[row][column] = float64(1)
	for i := 1; i <= k; i++ {
		record = make([][]float64, n)
		for i := 0; i < n; i++ {
			record[i] = make([]float64, n)
		}
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				for m := 0; m < len(positions); m++ {
					newr := r + positions[m][0]
					newc := c + positions[m][1]
					if newr < 0 || newc >= n || newr >= n || newc < 0 {
						continue
					}
					record[newr][newc] = record[newr][newc] + recordOld[r][c]/8
				}
			}
		}
		copy(recordOld, record)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans = ans + recordOld[i][j]
		}
	}
	return ans
}

// @lc code=end

