/*
 * @lc app=leetcode.cn id=1186 lang=golang
 *
 * [1186] 删除一次得到子数组最大和
 */

// @lc code=start
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maximumSum(arr []int) int {
	n := len(arr)
	var (
		ans         int
		record, sum []int
	)
	record = make([]int, n)
	sum = make([]int, n)
	sum[0] = arr[0]
	record[0] = arr[0]
	ans = arr[0]
	for i := 1; i < n; i++ {
		sum[i] = Max(sum[i-1]+arr[i], arr[i])
		record[i] = Max(sum[i-1], record[i-1]+arr[i])
		ans = Max(ans, Max(record[i], sum[i]))
	}
	return ans
}

// @lc code=end

