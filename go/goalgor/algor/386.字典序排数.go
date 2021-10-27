/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start
func lexicalOrder(n int) []int {
	var (
		ans     []int
		iterate func(int)
	)
	ans = []int{}
	iterate = func(x int) {
		if x > n {
			return
		}
		ans = append(ans, x)
		for i := x * 10; i < x*10+10; i++ {
			iterate(i)
		}
	}
	ans = []int{}
	for i := 1; i < 10; i++ {
		iterate(i)
	}
	return ans
}

// @lc code=end

