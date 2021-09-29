func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	record := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		record[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				record[i][j] = record[i-1][j-1] + 1
			} else {
				if record[i][j-1] < record[i-1][j] {
					record[i][j] = record[i-1][j]
				} else {
					record[i][j] = record[i][j-1]
				}
			}
		}
	}
	return record[m][n]
}