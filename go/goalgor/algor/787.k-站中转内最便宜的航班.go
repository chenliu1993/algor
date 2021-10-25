/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 */

// @lc code=start
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	record := make([][]int, n)
	for i := 0; i < n; i++ {
		record[i] = make([]int, k+2)
	}
	var iterate func(int, int) int
	iterate = func(start, transfers int) int {
		if transfers > k+1 {
			return 1000007
		}
		if start == dst {
			return 0
		}
		if record[start][transfers] != 0 {
			return record[start][transfers]
		}
		minPrice := 1000007
		for i := 0; i < len(flights); i++ {
			if flights[i][0] == start {
				minPrice = Min(minPrice, flights[i][2]+iterate(flights[i][1], transfers+1))
			}
		}
		record[start][transfers] = minPrice
		return minPrice
	}

	ans := iterate(src, 0)
	if ans == 1000007 {
		return -1
	}
	return ans
}

// @lc code=end

