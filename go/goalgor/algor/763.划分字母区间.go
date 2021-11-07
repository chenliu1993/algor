/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(s string) []int {
	n := len(s)
	var (
		curRight    int
		curLeft     int
		rightBorder map[byte]int = map[byte]int{}
		ans         []int
	)
	// Get each byte's range
	for i := 0; i < n; i++ {
		rightBorder[s[i]] = i
	}
	// Initialize curRight to first rightBorder
	curRight = rightBorder[s[0]]
	curLeft = 0
	for i := 1; i < n; i++ {
		if i > curRight {
			ans = append(ans, curRight-curLeft+1)
			curLeft = i
			curRight = rightBorder[s[i]]
		}
		if rightBorder[s[i]] > curRight {
			curRight = rightBorder[s[i]]
		}
	}
	ans = append(ans, len(s)-curLeft)
	return ans
}

// @lc code=end

