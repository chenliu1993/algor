/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉 K 位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	var (
		n   int
		stk []byte
	)
	n = len(num)
	if k >= n {
		return "0"
	}
	stk = []byte{}
	for i := 0; i < n; i++ {
		for len(stk) != 0 && stk[len(stk)-1] > num[i] {
			if k <= 0 {
				break
			}
			stk = stk[:len(stk)-1]
			k--
		}
		stk = append(stk, num[i])
	}
	if k != 0 {
		stk = stk[:len(stk)-k]
	}
	ans := strings.TrimLeft(string(stk), "0")
	if strings.Compare(ans, "") == 0 {
		return "0"
	}
	return ans
}

// @lc code=end

