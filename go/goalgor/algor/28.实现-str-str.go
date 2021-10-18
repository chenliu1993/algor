/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if m == 0 {
		return 0
	}
	var (
		i        int = 0
		jp       int = 0
		next     []int
		findNext func(string)
	)
	findNext = func(s string) {
		k := -1
		j := 0
		for j < m-1 {
			if k == -1 || s[j] == s[k] {
				j = j + 1
				k = k + 1
				next[j] = k
			} else {
				k = next[k]
			}
		}
	}
	next = make([]int, m)
	next[0] = -1
	findNext(needle)

	for i < n && jp < m {
		if jp == -1 || haystack[i] == needle[jp] {
			i = i + 1
			jp = jp + 1
		} else {
			jp = next[jp]
		}
	}
	if jp == m {
		return i - jp
	}
	return -1
}

// @lc code=end

