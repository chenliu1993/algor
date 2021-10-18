/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxProfit(prices []int) int {
	n := len(prices)
	record := make([][]int, n)
	for i := 0; i < n; i++ {
		record[i] = make([]int, 2)
	}
	record[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		record[i][0] = Max(record[i-1][0], record[i-1][1]+prices[i])
		record[i][1] = Max(record[i-1][0]-prices[i], record[i-1][1])
	}
	return record[n-1][0]
}

// @lc code=end

