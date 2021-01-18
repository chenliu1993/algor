package algorithm

func partition(s string) [][]string {
	var (
		n     = len(s)
		ans   [][]string
		check = func(s string) bool {
			n := len(s)
			for i := 0; i < n; i++ {
				if s[i] != s[n-1-i] {
					return false
				}
			}
			return true
		}
		dfs func(start int, path []string)
	)
	dfs = func(start int, path []string) {
		if start == n {
			ans = append(ans, append([]string{}, path...))
			return
		}
		for i := start; i < n; i++ {
			if !check(s[start : i+1]) {
				continue
			}
			path = append(path, s[start:i+1])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []string{})
	return ans
}
