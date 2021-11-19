/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start

func diffWaysToCompute(expression string) []int {
	var (
		divide func(string) []int
	)
	// `expression`` must be valid
	divide = func(s string) []int {
		if !strings.ContainsAny(s, "+-*") {
			res, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return []int{res}
		}
		ans := []int{}
		for i := 0; i < len(s); i++ {
			if s[i] == '+' || s[i] == '-' || s[i] == '*' {
				left := divide(s[:i])
				right := divide(s[i+1:])
				for k := 0; k < len(left); k++ {
					for j := 0; j < len(right); j++ {
						var temp int
						if s[i] == '+' {
							temp = left[k] + right[j]
						} else if s[i] == '-' {
							temp = left[k] - right[j]
						} else {
							temp = left[k] * right[j]
						}
						ans = append(ans, temp)
					}
				}
			}
		}
		return ans
	}
	return divide(expression)
}

// @lc code=end

