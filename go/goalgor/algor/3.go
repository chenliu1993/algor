func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return n
	}
	maxLen := 0
	rear := -1
	records := make(map[byte]int)
	for front := 0; front < n; front++ {
		if front != 0 {
			delete(records, s[front-1])
		}
		for rear+1 < n && records[s[rear+1]] == 0 {
			records[s[rear+1]] = 1
			rear = rear + 1
		}
		if maxLen < rear-front+1 {
			maxLen = rear - front + 1
		}
	}
	return maxLen
}