/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start

func generate(numRows int) [][]int {
	var (
		ans [][]int
	)
	ans = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
	}
	// First row only has one element ans it is surely one
	ans[0][0] = 1
	// Begin calculate
	for i := 1; i < numRows; i++ {
		for j := 0; j < len(ans[i]); j++ {
			if j == 0 {
				ans[i][j] = ans[i-1][j]
			} else if j == i {
				ans[i][j] = ans[i-1][j-1]
			} else {
				ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
			}
		}
	}
	return ans
}

func getRow(rowIndex int) []int {
	return generate(rowIndex + 1)[rowIndex]
}

// @lc code=end

