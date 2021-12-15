/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start
func maximumSwap(num int) int {
	var (
		j     int
		ans   int
		str   []string
		first bool
		stk   []int
	)
	first = true
	str = strings.Split(strconv.Itoa(num), "")
	stk = []int{}
	for j = 0; j < len(str); j++ {
		for len(stk) != 0 && str[stk[len(stk)-1]] < str[j] {
			first = false
			stk = stk[:len(stk)-1]
		}
		if !first {
			break
		}
		stk = append(stk, j)

	}
	if j == len(str) {
		return num
	}
	spliter := j
	maxEle := str[j]
	max := j
	for ; j < len(str); j++ {
		if str[j] >= maxEle {
			max = j
			maxEle = str[j]
		}
	}
	min := 0
	for i := 0; i < spliter; i++ {
		if str[i] < maxEle {
			min = i
			break
		}
	}
	str[min], str[max] = str[max], str[min]
	ans, _ = strconv.Atoi(strings.Join(str, ""))
	return ans
}

// @lc code=end

