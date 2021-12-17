/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen(target int, nums []int) int {
	var (
		left, right int
		sum         int64
		minLen      int
	)
	n := len(nums)
	sum = int64(0)
	left = 0
	right = 0
	minLen = len(nums)
	for right < n {
		sum = sum + int64(nums[right])
		for left <= right && sum >= int64(target) {
			minLen = Min(minLen, right-left+1)
			sum = sum - int64(nums[left])
			left++
		}
		right++
		if (right-left) == n && sum < int64(target) {
			minLen = 0
		}
	}
	return minLen
}

// @lc code=end

