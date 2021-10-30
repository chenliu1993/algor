/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

// @lc code=start

func removeDuplicateLetters(s string) string {
	var (
		stk    []byte
		left   []int
		record []bool
		n      int
	)
	n = len(s)
	left = make([]int, 26)
	for i := 0; i < n; i++ {
		left[s[i]-'a']++
	}
	record = make([]bool, 26)
	stk = []byte{}
	for i := 0; i < n; i++ {
		if !record[s[i]-'a'] {
			for len(stk) != 0 && stk[len(stk)-1] > s[i] {
				if left[stk[len(stk)-1]-'a'] == 0 {
					break
				}
				record[stk[len(stk)-1]-'a'] = false
				stk = stk[:len(stk)-1]

			}
			stk = append(stk, s[i])
			record[s[i]-'a'] = true
		}
		left[s[i]-'a']--
	}
	return string(stk)
}

// @lc code=end

