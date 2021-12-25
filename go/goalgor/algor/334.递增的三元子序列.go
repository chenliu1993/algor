/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	var (
		flag             bool
		stk, left, right []int
	)
	stk, left, right = []int{}, []int{}, []int{}
	for i := 0; i < n; i++ {
		right = append(right, n)
		left = append(left, -1)
	}
	for i := 0; i < n; i++ {
		for len(stk) != 0 && nums[stk[len(stk)-1]] < nums[i] {
			right[stk[len(stk)-1]] = i
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	stk = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stk) != 0 && nums[stk[len(stk)-1]] > nums[i] {
			left[stk[len(stk)-1]] = i
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	for i := 0; i < n; i++ {
		if left[i] != -1 && right[i] != n {
			flag = true
			break
		}
	}
	return flag
}

// @lc code=end

