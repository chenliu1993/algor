/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
func monotoneIncreasingDigits(n int) int {
	var (
		ans int
	)
	nStr := []byte(strconv.Itoa(n))
	i := 1
	for i < len(nStr) && nStr[i] >= nStr[i-1] {
		i++
	}
	if i < len(nStr) {
		for i > 0 && nStr[i] < nStr[i-1] {
			nStr[i-1]--
			i--
		}
		for i++; i < len(nStr); i++ {
			nStr[i] = '9'
		}
	}
	ans, _ = strconv.Atoi(string(nStr))
	return ans
}

// @lc code=end

// func monotoneIncreasingDigits(n int) int {
// 	var (
// 		ans                int
// 		isIncreasingDigits func(int) bool
// 	)
// 	isIncreasingDigits = func(input int) bool {
// 		inputStr := strconv.Itoa(input)
// 		stk := []int{}
// 		flag := true
// 		for i := len(inputStr) - 1; i >= 0; i-- {
// 			for len(stk) != 0 && inputStr[stk[len(stk)-1]]-'0' < inputStr[i]-'0' {
// 				flag = false
// 				stk = stk[:len(stk)-1]
// 			}
// 			stk = append(stk, i)
// 		}
// 		return flag
// 	}
// 	ans = 0
// 	for i := n; i >= 0; i-- {
// 		if isIncreasingDigits(i) {
// 			ans = i
// 			break
// 		}
// 	}
// 	return ans
// }

