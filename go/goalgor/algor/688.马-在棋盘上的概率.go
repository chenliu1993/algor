/*
 * @lc app=leetcode.cn id=688 lang=golang
 *
 * [688] “马”在棋盘上的概率
 */

// @lc code=start
var positions = [][]int{{-1, -2}, {1, -2}, {-1, 2}, {1, 2}, {-2, -1}, {-2, 1}, {2, -1}, {2, 1}}

func knightProbability(n int, k int, row int, column int) float64 {
	var (
		newi, newj int
		record     [][][]float64
	)

	k = k + 1
	record = make([][][]float64, k)
	for i := 0; i < k; i++ {
		record[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			record[i][j] = make([]float64, n)
		}
	}

	record[k-1][row][column] = float64(1)
	for k = k - 2; k >= 0; k-- {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for p := 0; p < len(positions); p++ {
					newi = i + positions[p][0]
					newj = j + positions[p][1]
					if newi < 0 || newi >= n || newj < 0 || newj >= n {
						continue
					}
					record[k][newi][newj] = record[k][newi][newj] + record[k+1][i][j]/float64(8)
				}
			}
		}
	}
	ans := float64(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans = ans + record[0][i][j]
		}
	}
	return ans
}

// @lc code=end

