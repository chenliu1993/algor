/*
 * @lc app=leetcode.cn id=712 lang=golang
 *
 * [712] 两个字符串的最小ASCII删除和
 */

// @lc code=start

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
func minimumDeleteSum(s1 string, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)
	var (
		sum    int
		record [][]int
	)
	sum = 0
	record = make([][]int, n1)
	for i := 0; i < n1; i++ {
		record[i] = make([]int, n2)
		if i != 0 {
			record[i][0] = record[i-1][0]
		}
		if s1[i] == s2[0] {
			record[i][0] = 2 * int(s2[0])
		}
		sum = sum + int(s1[i])
	}
	for i := 0; i < n2; i++ {
		if i != 0 {
			record[0][i] = record[0][i-1]
		}
		if s1[0] == s2[i] {
			record[0][i] = 2 * int(s1[0])
		}
		sum = sum + int(s2[i])
	}
	for i := 1; i < n1; i++ {
		for j := 1; j < n2; j++ {
			record[i][j] = Max(record[i-1][j], record[i][j-1])
			if s1[i] == s2[j] {
				record[i][j] = Max(record[i][j], record[i-1][j-1]+int(s1[i])+int(s2[j]))
			}
		}
	}
	return sum - record[n1-1][n2-1]
}

// @lc code=end

