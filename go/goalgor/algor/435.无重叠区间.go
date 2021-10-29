/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start

func eraseOverlapIntervals(intervals [][]int) int {
	var (
		count = 0
		last  []int
	)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	last = intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < last[1] {
			count++
			continue
		}
		last = intervals[i]
	}
	return count
}

// @lc code=end

