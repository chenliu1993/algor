/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */

// @lc code=start

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(x, y int) bool { return pairs[x][1] < pairs[y][1] })
	last := pairs[0]
	count := 1
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > last[1] {
			count++
			last = pairs[i]
			continue
		}
	}
	return count
}

// @lc code=end

