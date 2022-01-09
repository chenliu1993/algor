/*
 * @lc app=leetcode.cn id=807 lang=golang
 *
 * [807] 保持城市天际线
 */

// @lc code=start
func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var (
		ans         int
		west, south []int
	)
	west, south = make([]int, m), make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			south[i] = Max(south[i], grid[j][i])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			west[i] = Max(west[i], grid[i][j])
		}
	}

	fmt.Println(west, south)
	ans = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := Min(west[i], south[j])
			ans += cur - grid[i][j]
		}
	}
	return ans
}

// @lc code=end

