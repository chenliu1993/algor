/*
 * @lc app=leetcode.cn id=918 lang=golang
 *
 * [918] 环形子数组的最大和
 */

// @lc code=start
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	var (
		ans                         int
		sum                         int
		minSums1, minSums2, maxSums []int
	)
	maxSums = make([]int, n)
	minSums1 = make([]int, n)
	minSums2 = make([]int, n)
	sum = 0
	maxSums[0] = nums[0]
	minSums1[0] = nums[0]
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
	}
	ans = -30001
	for i := 1; i < n; i++ {
		maxSums[i] = Max(maxSums[i-1]+nums[i], nums[i])
		ans = Max(ans, maxSums[i])
		if i < n-1 {
			minSums1[i] = Min(minSums1[i-1]+nums[i], nums[i])
			ans = Max(ans, sum-minSums1[i])
		}
	}
	minSums2[1] = nums[1]
	for i := 2; i < n; i++ {
		minSums2[i] = Min(minSums2[i-1]+nums[i], nums[i])
		ans = Max(ans, sum-minSums2[i])
	}
	return ans
}

// @lc code=end

