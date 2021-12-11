/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start
func triangleNumber(nums []int) int {
	n := len(nums)
	var (
		count int
		cur   int
	)
	sort.Ints(nums)
	count = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cur = nums[i] + nums[j]
			for k := j + 1; k < n; k++ {
				if nums[k] < cur {
					count++
				}
			}
		}
	}
	return count
}

// @lc code=end

