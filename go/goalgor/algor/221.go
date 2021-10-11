func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	maxLength := 0
	record := make([][]int, m)
	for i := 0; i < m; i++ {
		record[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					record[i][j] = 1
				} else {
					record[i][j] = Min(Min(record[i-1][j], record[i][j-1]), record[i-1][j-1]) + 1
				}
				maxLength = Max(record[i][j], maxLength)
			}
		}
	}
	return maxLength * maxLength
}
