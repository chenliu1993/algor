/*
 * @lc app=leetcode.cn id=1035 lang=golang
 *
 * [1035] 不相交的线
 */

// @lc code=start
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	var (
		record [][]int
	)

	record = make([][]int, m)
	for i := 0; i < m; i++ {
		record[i] = make([]int, n)
	}
	if nums1[0] == nums2[0] {
		record[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if nums1[i] == nums2[0] {
			record[i][0] = 1
		} else {
			record[i][0] = record[i-1][0]
		}
	}
	for i := 1; i < n; i++ {
		if nums1[0] == nums2[i] {
			record[0][i] = 1
		} else {
			record[0][i] = record[0][i-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if nums1[i] == nums2[j] {
				record[i][j] = Max(record[i-1][j-1]+1, Max(record[i-1][j], record[i][j-1]))
			} else {
				record[i][j] = Max(record[i-1][j], record[i][j-1])
			}
		}
	}
	return record[m-1][n-1]
}

// @lc code=end

