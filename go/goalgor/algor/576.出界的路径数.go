/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 */

// @lc code=start

const MOD int = 1e9 + 7

var positions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var (
		ans               int
		recordOld, record [][]int
	)
	recordOld = make([][]int, m)
	for i := 0; i < m; i++ {
		recordOld[i] = make([]int, n)
	}
	recordOld[startRow][startColumn] = 1
	for time := 1; time <= maxMove; time++ {
		record = make([][]int, m)
		for k := 0; k < m; k++ {
			record[k] = make([]int, n)
		}

		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				if recordOld[r][c] > 0 {
					for p := 0; p < len(positions); p++ {
						newr := r + positions[p][0]
						newc := c + positions[p][1]
						if newr < 0 || newr >= m || newc < 0 || newc >= n {
							ans = (ans + recordOld[r][c]) % MOD
						} else {
							record[newr][newc] = (record[newr][newc] + recordOld[r][c]) % MOD
						}
					}
				}

			}
		}
		copy(recordOld, record)
	}

	return ans
}

// @lc code=end

