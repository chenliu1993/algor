/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	n := len(s)
	var (
		ans    []string
		record [][]bool
	)
	record = make([][]bool, n)
	for i := 0; i < n; i++ {
		record[i] = make([]bool, n)
		record[i][i] = true
	}

	ans = strings.Split(s, "")
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && (j == i+1 || record[i+1][j-1]) {
				record[i][j] = true
				ans = append(ans, s[i:j+1])
			}
		}
	}

	return len(ans)
}

// @lc code=end

