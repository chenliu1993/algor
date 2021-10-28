/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

// @lc code=start

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func integerReplacement(n int) int {
	var (
		record func(int) int
	)
	record = func(x int) int {
		if x == 1 {
			return 0
		}
		if x%2 == 0 {
			return 1 + record(x/2)
		}
		return 1 + Min(record(x-1), record(x+1))
	}

	return record(n)
}

// func integerReplacement(n int) int {
// 	var (
// 		record []int
// 	)
// 	record = make([]int, n+1)
// 	// start from n = 2
// 	for i := 2; i < n+1; i++ {
// 		if i%2 == 0 {
// 			record[i] = 1 + record[i/2]
// 		} else {
// 			record[i] = 2 + Min(record[(i-1)/2], record[(i+1)/2])
// 		}
// 	}
// 	return record[n]
// }

// @lc code=end

