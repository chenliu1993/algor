/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}
	var first, rear int
	first = 0
	rear = 0

	for rear < len(nums) {
		for nums[first] != 0 {
			first++
			if first == len(nums) {
				return
			}
		}
		for nums[rear] == 0 {
			rear++
			if rear == len(nums) {
				return
			}
		}

		if rear < first {
			rear = first
		}
		temp := nums[rear]
		nums[rear] = nums[first]
		nums[first] = temp
	}
}

// @lc code=end

