/*
 * @lc app=leetcode.cn id=1130 lang=golang
 *
 * [1130] 叶值的最小代价生成树
 */

// @lc code=start
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func mctFromLeafValues(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}

	var (
		biggest [][]int
		record  [][]int
	)
	record = make([][]int, n)
	biggest = make([][]int, n)
	for i := 0; i < n; i++ {
		record[i] = make([]int, n)
		for j := 0; j < n; j++ {
			record[i][j] = 2147483647
		}
		record[i][i] = 0
		biggest[i] = make([]int, n)
		biggest[i][i] = arr[i]
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			biggest[i][j] = Max(biggest[i][j-1], arr[j])
		}
	}
	for i := 0; i < n-1; i++ {
		record[i][i+1] = arr[i] * arr[i+1]
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			for k := i; k < j; k++ {
				record[i][j] = Min(record[i][j], record[i][k]+record[k+1][j]+biggest[i][k]*biggest[k+1][j])
			}
		}
	}
	return record[0][n-1]
}

// @lc code=end

