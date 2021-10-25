func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	weights := make([][]int, n)
	for i := 0; i < len(weights); i++ {
		weights[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			weights[i][j] = 10001
		}
	}
	for i := 0; i < len(flights); i++ {
		weights[flights[i][0]][flights[i][1]] = flights[i][2]
	}
	weights[src][src] = 0
	curPrice := 0
	minPrice := 10001
	visited := make([]int, n)
	visited[src] = 1
	var iterate func(int, int)
	iterate = func(start, transfers int) {
		if start == dst {
			if transfers <= k+1 {
				minPrice = Min(minPrice, curPrice)
			}
			return
		}
		for i := 0; i < len(weights[start]); i++ {
			if i == start || visited[i] == 1 {
				continue
			}
			curPrice = curPrice + weights[start][i]
			visited[i] = 1
			iterate(i, transfers+1)
			visited[i] = 0
			curPrice = curPrice - weights[start][i]
		}
	}
	iterate(src, 0)
	if minPrice == 10001 {
		return -1
	}
	return minPrice
}

// over time limit