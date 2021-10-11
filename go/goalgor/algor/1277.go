func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	count := 0
	record := make([][]int, m)
	for i := 0; i < m; i++ {
		record[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				record[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				record[i][j] = 0
			} else {
				record[i][j] = Min(Min(record[i-1][j], record[i][j-1]), record[i-1][j-1]) + 1
			}
			count = count + record[i][j]
		}
	}
	return count
}