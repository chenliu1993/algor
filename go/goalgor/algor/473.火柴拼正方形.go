/*
 * @lc app=leetcode.cn id=473 lang=golang
 *
 * [473] 火柴拼正方形
 */

// @lc code=start
func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	var (
		sum         int64
		stickLength int64
		ans         bool
		bucket      []int64
		dfs         func(int) bool
	)
	for i := 0; i < n; i++ {
		sum = sum + int64(matchsticks[i])
	}
	if sum%4 != 0 {
		return false
	}
	stickLength = sum / 4
	sort.Slice(matchsticks, func(i, j int) bool {
		return i > j
	})
	bucket = make([]int64, 4)
	dfs = func(start int) bool {
		if ans {
			return true
		}
		if start == n {
			return bucket[0] == bucket[1] && bucket[1] == bucket[2] && bucket[2] == bucket[3]
		}
		if int64(matchsticks[start]) > stickLength {
			return false
		}

		for i := 0; i < 4; i++ {
			temp := bucket[i]
			if temp+int64(matchsticks[start]) <= stickLength || (i > 0 && bucket[i] == bucket[i-1]) {
				bucket[i] = temp + int64(matchsticks[start])
				if dfs(start + 1) {
					return true
				}
				bucket[i] = temp
			}
		}
		return false
	}
	return dfs(0)
}

// @lc code=end

