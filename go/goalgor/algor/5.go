func longestPalindrome(s string) string {
	var ans string
	n := len(s)
	record := make([][]int, n)
	for i := 0; i < n; i++ {
		record[i] = make([]int, n)
		record[i][i] = 1
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := l + i
			if l == 0 {
				record[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					record[i][j] = 1
				}
			} else {
				if s[i] == s[j] && record[i+1][j-1] == 1 {
					record[i][j] = 1
				}
			}
			if record[i][j] == 1 && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
