/*
 * @lc app=leetcode.cn id=526 lang=golang
 *
 * [526] 优美的排列
 */

// @lc code=start
func countArrangement(n int) int {
	var (
		curSlice []int
		visited  []int
		ans      [][]int
		iterate  func(int)
	)
	iterate = func(posi int) {
		if posi == n+1 {
			tempSlice := make([]int, len(curSlice))
			copy(tempSlice, curSlice)
			ans = append(ans, tempSlice)
			return
		}
		for i := 1; i < len(visited); i++ {
			if visited[i] != 1 {
				if posi%i == 0 || i%posi == 0 {
					visited[i] = 1
					curSlice = append(curSlice, i)
					iterate(posi + 1)
					curSlice = curSlice[:len(curSlice)-1]
					visited[i] = 0
				}
			}
		}
	}
	ans = [][]int{}
	visited = make([]int, n+1)
	curSlice = []int{}
	iterate(1)
	return len(ans)
}

// @lc code=end

