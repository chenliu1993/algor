/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func isPali(s string) bool {
	start, end := 0, len(s)-1
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func partition(s string) [][]string {
	n := len(s)
	var (
		ans     [][]string
		iterate func(int, []string)
	)
	ans = [][]string{}
	iterate = func(pos int, path []string) {
		if pos >= n {
			ans = append(ans, append([]string{}, path...))
			return
		}
		for i := pos; i < n; i++ {
			if !isPali(s[pos : i+1]) {
				continue
			}
			path = append(path, s[pos:i+1])
			iterate(i+1, path)
			path = path[:len(path)-1]
		}
	}
	iterate(0, []string{})
	return ans

}

// @lc code=end

