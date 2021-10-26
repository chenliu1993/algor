/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func integerBreak(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	record := make([]int, n+1)
	for i := 2; i <= n; i++ {
		cur := 0
		for j := 1; j < i; j++ {
			cur = Max(Max(cur, j*(i-j)), j*record[i-j])
		}
		record[i] = cur
	}
	return record[n]
}

// @lc code=end

