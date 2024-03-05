/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	base := make([]int, len(nums))
	base[0] = nums[0]
	count := 0
	for i := 1; i < len(nums); i++ {
		base[i] = base[i-1] + nums[i]

	}

	for i := 0; i < len(base); i++ {
		if base[i] == k {
			count++
		}
	}

	for i := len(base) - 1; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if k == base[i]-base[j] {
				count++
			}
		}
	}

	return count
}

// @lc code=end

