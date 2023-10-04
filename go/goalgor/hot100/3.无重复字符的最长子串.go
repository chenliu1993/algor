/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	maxlen := 1
	record := make(map[byte]bool)
	left := -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(record, s[i-1])
		}
		for left+1 < len(s) && !record[s[left+1]] {
			record[s[left+1]] = true
			left++
		}
		maxlen = max(maxlen, left-i+1)
	}
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

