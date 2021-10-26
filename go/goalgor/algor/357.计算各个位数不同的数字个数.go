/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 */

// @lc code=start

func Pow10(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	if n%2 == 0 {
		return Pow10(n/2) * Pow10(n/2)
	}
	return Pow10((n-1)/2) * Pow10((n-1)/2) * 10
}

func countNumbersWithUniqueDigits(n int) int {
	var (
		count     int
		bits      []int
		dict      map[int]int
		rightMost int = Pow10(n)
		iterate   func(int)
	)
	// 0 marks unused
	// 1 marks used
	bits = make([]int, 10)
	dict = map[int]int{}
	iterate = func(x int) {
		if x >= rightMost {
			return
		}
		if _, ok := dict[x]; !ok {
			count = count + 1
			dict[x] = 1
		}
		for i := 0; i <= 9; i++ {
			if bits[i] == 0 {
				temp := x*10 + i
				bits[i] = 1
				iterate(temp)
				bits[i] = 0
			}
		}
	}
	count = 0
	iterate(0)
	return count
}

// @lc code=end

