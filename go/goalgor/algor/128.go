func maxProfit(prices []int) int {
	n := len(prices)
	maxPro := 0
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if maxPro < prices[i]-minPrice {
			maxPro = prices[i] - minPrice
		}
	}
	return maxPro
}