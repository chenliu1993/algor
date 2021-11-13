/*
 * @lc app=leetcode.cn id=1785 lang=golang
 *
 * [1785] 构成特定和需要添加的最少元素
 */

// @lc code=start
func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func minElements(nums []int, limit int, goal int) int {
	var (
		sum int
		sub int
	)
	n := len(nums)
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
	}
	carry := 0
	sub = Abs(goal - sum)
	if sub%limit != 0 {
		carry = 1
	}
	return carry + sub/limit
}

// @lc code=end

