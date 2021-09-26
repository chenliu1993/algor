func firstUniqChar(s string) int {
	n := len(s)
	dict := make(map[byte]int, n)
	for i := 0; i < n; i++ {
		if _, ok := dict[s[i]]; !ok {
			dict[s[i]] = 1
		} else {
			dict[s[i]] = dict[s[i]] + 1
		}
	}
	for i := 0; i < n; i++ {
		if dict[s[i]] == 1 {
			return i
		}
	}
	return -1
}