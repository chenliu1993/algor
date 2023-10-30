/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}
	for fixed := 0; fixed < len(nums); fixed++ {
		if nums[fixed] > 0 {
			return res
		}
		if fixed > 0 && nums[fixed] == nums[fixed-1] {
			continue
		}
		left := fixed + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[fixed]
			if sum == 0 {
				res = append(res, []int{nums[fixed], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}

	}
	return res
}

// @lc code=end

