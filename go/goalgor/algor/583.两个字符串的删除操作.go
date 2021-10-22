/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */

// @lc code=start

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	record := make([][]int, n)
	for i := 0; i < n; i++ {
		record[i] = make([]int, m)
	}
	if word1[0] == word2[0] {
		record[0][0] = 1
	}
	for i := 1; i < n; i++ {
		if word2[0] == word1[i] {
			record[i][0] = 1
		} else {
			record[i][0] = record[i-1][0]
		}
	}
	for i := 1; i < m; i++ {
		if word1[0] == word2[i] {
			record[0][i] = 1
		} else {
			record[0][i] = record[0][i-1]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if word1[i] == word2[j] {
				record[i][j] = record[i-1][j-1] + 1
			} else {
				record[i][j] = Max(record[i-1][j], record[i][j-1])
			}
		}
	}
	return n + m - 2*record[n-1][m-1]
}

// @lc code=end

