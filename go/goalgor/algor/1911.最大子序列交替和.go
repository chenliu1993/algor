/*
 * @lc app=leetcode.cn id=1911 lang=golang
 *
 * [1911] 最大子序列交替和
 */

// @lc code=start
func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	even := make([]int64, n)
	odd := make([]int64, n)
	odd[0] = int64(nums[0])
	for i := 1; i < n; i++ {
		even[i] = Max(even[i-1], odd[i-1]-int64(nums[i]))
		odd[i] = Max(odd[i-1], even[i-1]+int64(nums[i]))
	}
	fmt.Println(even)
	fmt.Println(odd)
	return Max(even[n-1], odd[n-1])
}

// @lc code=end

