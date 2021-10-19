/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start

func Pow2(time int) int {
	if time == 0 {
		return 1
	}
	if time%2 == 0 {
		return Pow2(time/2) * Pow2(time/2)
	}
	return 2 * Pow2(time-1)
}

func getBinary(num int) int {
	var (
		left int
		ans  []int = []int{}
		sum  int
	)
	for num != 0 {
		if num%2 == 1 {
			left = 0
		} else {
			left = 1
		}
		ans = append(ans, left)
		num = num / 2
	}
	sum = 0
	for i := len(ans) - 1; i >= 0; i-- {
		sum = sum + ans[i]*Pow2(i)
		fmt.Println(ans[i], Pow2(i))
	}
	return sum
}

func findComplement(num int) int {
	return getBinary(num)
}

// func findComplement(num int) int {
//     return num ^ (1<<bits.Len32(uint32(num)) - 1)
// }

// @lc code=end

