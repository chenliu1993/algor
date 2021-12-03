/*
 * @lc app=leetcode.cn id=769 lang=golang
 *
 * [769] 最多能完成排序的块
 */

// @lc code=start
func maxChunksToSorted(arr []int) int {
	n := len(arr)
	var (
		blocks int
		met    map[int]bool
	)
	blocks = 0
	met = map[int]bool{}
	for i := 0; i < n; i++ {
		// i and arr[i] is grouped
		if _, ok := met[i]; !ok {
			met[i] = true
		} else {
			delete(met, i)
		}
		if _, ok := met[arr[i]]; !ok {
			met[arr[i]] = true
		} else {
			delete(met, arr[i])
		}
		if len(met) == 0 {
			blocks++
		}
	}
	return blocks
}

// @lc code=end

