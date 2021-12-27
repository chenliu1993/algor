/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start

func firstMissingPositive(nums []int) int {
	n := len(nums)
	hash := func(x int) int {
		return x - 1
	}
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[hash(nums[i])] != nums[i] {
			nums[hash(nums[i])], nums[i] = nums[i], nums[hash(nums[i])]
		}
	}
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return n + 1
}

// @lc code=end

