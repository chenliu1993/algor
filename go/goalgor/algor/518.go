func change(amount int, coins []int) int {
	n := len(coins)
	record := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		record[i] = make([]int, amount+1)
	}
	record[0][0] = 1
	for i := 1; i <= n; i++ {
		v := coins[i-1]
		for j := 0; j <= amount; j++ {
			for k := 0; k*v <= j; k++ {
				record[i][j] = record[i][j] + record[i-1][j-k*v]
			}
		}
	}
	return record[n][amount]
}