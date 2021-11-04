/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	n := len(nums)
	tempNums := make([]int, n)
	copy(tempNums, nums)
	sort.Ints(tempNums)
	left, right := 0, len(nums)-1
	for nums[left] == tempNums[left] {
		left++
	}
	for nums[right] == tempNums[right] {
		right--
	}
	return right - left + 1
}

// @lc code=end

