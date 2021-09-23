func canCompleteCircuit(gas []int, cost []int) int {
	var j int
	n := len(gas)
	can := -1
	for i := 0; i < n; i++ {
		sum := gas[i]
		// This is the start
		if i == n-1 {
			j = 0
		} else {
			j = i + 1
		}
		for j != i && j < n {
			if j != 0 {
				sum = sum - cost[j-1]
			} else {
				sum = sum - cost[n-1]
			}
			if sum < 0 {
				break
			}
			sum = sum + gas[j]
			j++
			j = j % n
		}
		if sum >= 0 && j == i {
			if j != 0 && sum-cost[j-1] >= 0 {
				can = i
			} else if j == 0 && sum-cost[n-1] >= 0 {
				can = i
			}
		}
	}
	return can
}