func longestPalindrome(s string) int {
	length := 0
	dict := map[string]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := dict[string(s[i])]; !ok {
			dict[string(s[i])] = 1
		} else {
			dict[string(s[i])] = dict[string(s[i])] + 1
		}
	}
	for _, v := range dict {
		length = length + int(v/2)*2
		if v%2 == 1 && length%2 == 0 {
			length = length + 1
		}
	}
	return length
}