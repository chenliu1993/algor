/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp1, dp2 := make([]int, len(nums)), make([]int, len(nums))
	dp1[0] = nums[0] // last one is not tolen
	dp2[0] = 0       // last one is stolen
	dp1[1] = nums[0]
	dp2[1] = nums[1]
	for i := 2; i < len(nums)-1; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])

	}

	for i := 2; i < len(nums); i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
	}
	return max(dp1[len(nums)-2], dp2[len(nums)-1])
}

// @lc code=end

