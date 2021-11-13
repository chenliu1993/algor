/*
 * @lc app=leetcode.cn id=1921 lang=golang
 *
 * [1921] 消灭怪物的最大数量
 */

// @lc code=start
func eliminateMaximum(dist []int, speed []int) int {
	var (
		times []int
	)
	times = make([]int, len(speed))
	for i := 0; i < len(speed); i++ {
		times[i] = (dist[i] - 1) / speed[i]
	}
	sort.Ints(times)
	for i := 0; i < len(speed); i++ {
		if times[i] < i {
			return i
		}
	}
	return len(speed)
}

// @lc code=end

