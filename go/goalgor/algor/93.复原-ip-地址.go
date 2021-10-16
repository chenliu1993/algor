/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func validBit(s string) bool {
	n := len(s)
	if n > 1 {
		if s[0] == '0' {
			return false
		}
	}
	sval, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if sval >= 0 && sval <= 255 {
		return true
	}
	return false
}

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 {
		return []string{}
	}
	var (
		count  int
		curStr string
		ans    []string
	)
	var iterate func(string)
	count = 0
	iterate = func(str string) {
		if count == 3 {
			if len(str) <= 3 && len(str) > 0 && validBit(str) {
				ans = append(ans, curStr+str)
			}
			return
		}

		tempStr := curStr

		count = count + 1
		if len(str) >= 1 && validBit(str[:1]) {
			curStr = curStr + str[:1] + "."
			iterate(str[1:])
			curStr = tempStr
		}

		if len(str) >= 2 && validBit(str[:2]) {
			curStr = curStr + str[:2] + "."
			iterate(str[2:])
			curStr = tempStr
		}

		if len(str) >= 3 && validBit(str[:3]) {
			curStr = curStr + str[:3] + "."
			iterate(str[3:])
			curStr = tempStr
		}
		count = count - 1
	}

	iterate(s)

	return ans
}

// @lc code=end

