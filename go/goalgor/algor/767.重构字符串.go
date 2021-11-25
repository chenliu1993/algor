/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 */

// @lc code=start
func reorganizeString(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	var (
		ans    string
		holder byte
	)

	record := map[byte]int{}
	left := []byte{}
	for i := 0; i < n; i++ {
		if _, ok := record[s[i]]; !ok {
			record[s[i]] = 1
			left = append(left, s[i])
		} else {
			record[s[i]]++
		}
	}
	sort.Slice(left, func(i, j int) bool {
		return record[left[i]] > record[left[j]]
	})
	holder = left[0]
	ans = string(holder)
	record[holder] = record[holder] - 1
	conCount := 0
	for {
		conCount = 0
		sort.Slice(left, func(i, j int) bool {
			return record[left[i]] > record[left[j]]
		})
		for i := 0; i < len(left); i++ {
			if holder == left[i] || record[left[i]] == 0 {
				conCount++
				continue
			}
			ans = ans + string(left[i])
			holder = left[i]
			record[left[i]] = record[left[i]] - 1
			break
		}
		if conCount == len(left) {
			break
		}
	}
	if len(ans) != n {
		ans = ""
	}
	return ans
}

// @lc code=end

