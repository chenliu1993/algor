/*
 * @lc app=leetcode.cn id=790 lang=golang
 *
 * [790] 多米诺和托米诺平铺
 */

// @lc code=start
func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	base := math.Pow10(9) + 7
	state := make([][]float64, n)
	for i := 0; i < n; i++ {
		state[i] = make([]float64, 3)
	}
	// Do a initialization
	state[0][0] = 0
	state[0][1] = 0
	state[0][2] = 1
	// Begin state find
	for i := 1; i < n; i++ {
		if i == 1 {
			state[i][0] = 1
			state[i][1] = 1
			state[i][2] = 2
		} else {
			state[i][0] = math.Mod(state[i-1][1]+state[i-2][2], base)
			state[i][1] = math.Mod(state[i-1][0]+state[i-2][2], base)
			state[i][2] = math.Mod(state[i-1][0]+state[i-1][1]+state[i-1][2]+state[i-2][2], base)
		}
	}
	return int(state[n-1][2])
}

// @lc code=end

