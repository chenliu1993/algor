func deckRevealedIncreasing(deck []int) []int {
	var stk []int
	arr := make([]int, len(deck))

	for idx, _ := range arr {
		arr[idx] = idx
	}

	for len(arr) > 0 {
		stk = append(stk, arr[0])
		arr = arr[1:]
		if len(arr) >= 1 {
			arr = append(arr, arr[0])
			arr = arr[1:]
		}
	}

	sort.Ints(deck)
	resDeck := make([]int, len(stk))
	for k, v := range stk {
		resDeck[v] = deck[k]
	}
	return resDeck
}