/*
 * @lc app=leetcode.cn id=672 lang=golang
 *
 * [672] 灯泡开关 Ⅱ
 */

// @lc code=start

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func flipLights(n int, presses int) int {
	n = Min(n, 3)
	if presses == 0 {
		return 1
	} else if presses == 1 {
		return n + 1
	} else if presses == 2 {
		if n == 1 {
			return 2
		} else if n == 2 {
			return 4
		} else if n == 3 {
			return 7
		}
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 4
	}
	return 8
}

// @lc code=end

