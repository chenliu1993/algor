func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func coinChange(coins []int, amount int) int {
	n := len(coins)
	MaxInt := 1000000
	record := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		record[i] = make([]int, amount+1)
	}

	for j := 0; j <= n; j++ {
		for i := 0; i < amount+1; i++ {
			record[j][i] = MaxInt
		}
	}
	record[0][0] = 0

	for i := 1; i < n+1; i++ {
		val := coins[i-1]
		for j := 0; j < amount+1; j++ {
			for k := 0; k*val <= j; k++ {
				record[i][j] = Min(record[i][j], record[i-1][j-k*val]+k)
			}
		}
	}
	if record[n][amount] == MaxInt {
		return -1
	}
	return record[n][amount]
}