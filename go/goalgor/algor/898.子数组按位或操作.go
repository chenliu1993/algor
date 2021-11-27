/*
 * @lc app=leetcode.cn id=898 lang=golang
 *
 * [898] 子数组按位或操作
 */

// @lc code=start

func subarrayBitwiseORs(arr []int) int {
	n := len(arr)
	var (
		sum, maxSum int64
		record      map[int64]int
	)
	record = map[int64]int{}
	for i := 0; i < n; i++ {
		sum = int64(arr[i])
		for j := i; j < n; j++ {
			sum = sum | int64(arr[j])
			if _, ok := record[sum]; !ok {
				record[sum] = 1
			}
			if sum == maxSum {
				break
			}
		}
		if i == 0 {
			maxSum = sum
		}
	}
	return len(record)
}

// @lc code=end

