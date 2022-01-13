/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
var dict = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
	4:    "IV",
	9:    "IX",
	40:   "XL",
	90:   "XC",
	400:  "CD",
	900:  "CM",
}

func getNum(n int) string {
	if n < 10 {
		if n == 4 || n == 9 {
			return dict[n]
		}
		if n < 5 {
			return strings.Repeat(dict[1], n)
		} else if n < 10 && n >= 5 {
			return dict[5] + strings.Repeat(dict[1], n-5)
		}
	} else if n >= 10 && n < 100 {
		if n == 40 || n == 90 || n == 10 || n == 50 {
			return dict[n]
		}
		if n < 50 {
			return strings.Repeat(dict[10], n/10)
		} else if n > 50 && n < 100 {
			return dict[50] + strings.Repeat(dict[10], (n-50)/10)
		}
	} else if n >= 100 && n < 1000 {
		if n == 100 || n == 400 || n == 900 || n == 500 {
			return dict[n]
		}
		if n < 500 {
			return strings.Repeat(dict[100], n/100)
		} else if n > 500 && n < 1000 {
			return dict[500] + strings.Repeat(dict[100], (n-500)/100)
		}
	}
	if n == 1000 {
		return dict[n]
	}
	return strings.Repeat(dict[1000], n/1000)
}

func intToRoman(num int) string {
	ans := ""
	bit := 1
	var carry int
	for num != 0 {
		carry = (num % 10) * bit
		ans = getNum(carry) + ans
		num = num / 10
		bit = bit * 10
	}
	return ans
}

// @lc code=end

