/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}
	var neg bool
	if s[0] == '-' {
		neg = true
		s = s[1:]
	} else if s[0] == '+' {
		neg = false
		s = s[1:]
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			s = s[:i]
			break
		}
	}
	var ans float64
	for i := 0; i < len(s); i++ {
		ans += float64(s[i]-'0') * math.Pow(10, float64(len(s)-1-i))
	}
	if neg {
		ans = ans * float64(-1)
	}
	if ans < math.MinInt32 {
		ans = math.MinInt32
	} else if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	return int(ans)
}

// @lc code=end

