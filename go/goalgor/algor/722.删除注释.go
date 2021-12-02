/*
 * @lc app=leetcode.cn id=722 lang=golang
 *
 * [722] 删除注释
 */

// @lc code=start
func removeComments(source []string) []string {
	sum := ""
	stk := ""
	ans := []string{}
	for i := 0; i < len(source); i++ {
		sum = sum + "##" + source[i]
	}
	sum = strings.TrimPrefix(sum, "##")
	sum = sum + "##"
	for i := 0; i < len(sum); i++ {
		if sum[i] == '/' {
			if sum[i+1] == '*' {
				i = strings.Index(sum[i+2:], "*/") + 3 + i
			} else if sum[i+1] == '/' {
				i = strings.Index(sum[i+2:], "##") + 1 + i
			} else {
				stk = stk + string(sum[i])
			}
		} else {
			stk = stk + string(sum[i])
		}
	}
	for strings.Contains(stk, "##") {
		idx := strings.Index(stk, "##")
		if idx == 0 {
			stk = stk[idx+2:]
			continue
		}
		ans = append(ans, stk[:idx])
		stk = stk[idx+2:]
	}
	if len(stk) != 0 {
		ans = append(ans, stk)
	}
	return ans
}

// @lc code=end

