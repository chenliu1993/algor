/*
 * @lc app=leetcode.cn id=1105 lang=golang
 *
 * [1105] 填充书架
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

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	if n == 1 {
		return books[0][1]
	}
	var (
		height int
		width  int
		record []int
	)
	record = make([]int, n)
	record[0] = books[0][1]
	for i := 1; i < len(books); i++ {
		height = books[i][1]
		width = books[i][0]
		record[i] = record[i-1] + books[i][1]
		for j := i - 1; j >= 0; j-- {
			height = Max(height, books[j][1])
			width = width + books[j][0]
			if width > shelfWidth {
				break
			}
			if j == 0 {
				record[i] = Min(record[i], height)
			} else {
				record[i] = Min(record[i], record[j-1]+height)
			}
		}
	}
	return record[n-1]
}

// @lc code=end

