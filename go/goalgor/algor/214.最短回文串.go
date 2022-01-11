/*
 * @lc app=leetcode.cn id=214 lang=golang
 *
 * [214] 最短回文串
 */

// @lc code=start
func shortestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	var (
		ans  string
		j, k int
		next []int
	)

	ans = ""
	for i := n - 1; i >= 0; i-- {
		ans += string(s[i])
	}
	ans = s + "#" + ans

	n = len(ans)
	next = make([]int, n)
	j, k = 0, -1
	next[0] = -1
	for j < n-1 {
		if k == -1 || ans[j] == ans[k] {
			k++
			j++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	maxLen := next[len(ans)-1] + 1
	ans = ""
	for i := maxLen; i < len(s); i++ {
		ans = string(s[i]) + ans
	}
	return ans + s
}

// @lc code=end

