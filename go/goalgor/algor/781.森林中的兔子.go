/*
 * @lc app=leetcode.cn id=781 lang=golang
 *
 * [781] 森林中的兔子
 */

// @lc code=start
func numRabbits(answers []int) int {
	n := len(answers)
	dict := map[int]int{}
	for i := 0; i < n; i++ {
		if _, ok := dict[answers[i]]; !ok {
			dict[answers[i]] = 1
		} else {
			dict[answers[i]]++
		}
	}
	for k, v := range dict {
		if k == 0 {
			continue
		}
		for {
			if k+1-v >= 0 {
				break
			}
			v -= (k + 1)
		}
		n += k + 1 - v

	}
	return n
}

// @lc code=end

