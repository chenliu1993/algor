/*
 * @lc app=leetcode.cn id=873 lang=golang
 *
 * [873] 最长的斐波那契子序列的长度
 */

// @lc code=start
func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	var (
		maxLen        int
		length        int
		first, second int
		res           int
		exists        map[int]bool
	)
	exists = map[int]bool{}
	for i := 0; i < n; i++ {
		exists[arr[i]] = true
	}
	maxLen = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			length = 2
			first = arr[i]
			second = arr[j]
			res = first + second
			for exists[res] {
				length++
				first = second
				second = res
				res = first + second
			}
			maxLen = Max(maxLen, length)
		}
	}
	if maxLen == 2 {
		return 0
	}
	return maxLen
}

// @lc code=end

// func lenLongestFibSubseq(arr []int) int {
// 	n := len(arr)
// 	var (
// 		maxLen int
// 		dict   map[int]int
// 		record [][]int
// 	)
// 	dict = map[int]int{}
// 	record = make([][]int, n)
// 	maxLen = 0
// 	for i := 0; i < n; i++ {
// 		dict[arr[i]] = i
// 		record[i] = make([]int, n)
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < i; j++ {
// 			k, ok := dict[arr[i]-arr[j]]
// 			if (arr[i]-arr[j] < arr[j]) && ok {
// 				record[j][i] = record[k][j] + 1
// 				maxLen = Max(maxLen, record[j][i]+2)
// 			}
// 		}
// 	}
// 	if maxLen < 3 {
// 		return 0
// 	}
// 	return maxLen
// }