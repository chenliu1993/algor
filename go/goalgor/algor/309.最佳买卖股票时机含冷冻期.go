/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
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
	coolDay := -1
	record := make([][]int, n)
	for i := 0; i < n; i++ {
		record[i] = make([]int, 2)
	}
	record[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		if record[i-1][0] < record[i-1][1]+prices[i] {
			record[i][0] = record[i-1][1] + prices[i]
			coolDay = i
		} else {
			record[i][0] = record[i-1][0]
		}
		if coolDay == i-1 {
			record[i][1] = Max(record[i-2][0]-prices[i], record[i-1][1])
		} else {
			record[i][1] = Max(record[i-1][0]-prices[i], record[i-1][1])
		}
	}
	return record[n-1][0]
}

// @lc code=end

