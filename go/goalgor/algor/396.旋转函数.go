/*
 * @lc app=leetcode.cn id=396 lang=golang
 *
 * [396] 旋转函数
 */

// @lc code=start
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxRotateFunction(nums []int) int {
	n := len(nums)
	var (
		posi   int
		ans    int
		sum    int
		rotate int
	)
	sum = 0
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
		// k=0, initialize the nums
		rotate = rotate + i*nums[i]
	}
	ans = rotate
	posi = n - 1
	for i := 1; i < n; i++ {
		rotate = rotate + sum - n*nums[posi]
		posi = posi - 1
		ans = Max(ans, rotate)
	}
	return ans
}

// @lc code=end

