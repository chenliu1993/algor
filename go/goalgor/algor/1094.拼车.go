/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 */

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	n := len(trips)
	sum := make([]int, 1001)
	for i := 0; i < n; i++ {
		for j := trips[i][1]; j < trips[i][2]; j++ {
			sum[j] += trips[i][0]
		}
	}
	sort.Ints(sum)
	if sum[len(sum)-1] > capacity {
		return false
	}
	return true
}

// @lc code=end

