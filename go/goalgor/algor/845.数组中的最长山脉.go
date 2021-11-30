/*
 * @lc app=leetcode.cn id=845 lang=golang
 *
 * [845] 数组中的最长山脉
 */

// @lc code=start

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func longestMountain(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}
	var (
		left     []int
		right    []int
		leftIdx  int
		rightIdx int
		maxLen   int
	)
	maxLen = 0
	for i := 1; i < n-1; i++ {
		left = []int{}
		leftIdx = -1
		rightIdx = -1
		right = []int{}
		for l := i; l >= 0; l-- {
			for len(left) != 0 && arr[left[len(left)-1]] <= arr[l] {
				if leftIdx == -1 {
					leftIdx = left[len(left)-1]
					l = 0
					break
				}
				left = left[:len(left)-1]
			}
			left = append(left, l)
		}
		if leftIdx == i {
			continue
		}
		if leftIdx == -1 {
			leftIdx = 0
		}

		for r := i; r < n; r++ {
			for len(right) != 0 && arr[right[len(right)-1]] <= arr[r] {
				if rightIdx == -1 {
					rightIdx = right[len(right)-1]
				}
				r = n - 1
				right = right[:len(right)-1]
				break
			}
			right = append(right, r)
		}
		if rightIdx == i {
			continue
		}
		if rightIdx == -1 {
			rightIdx = n - 1
		}
		maxLen = Max(maxLen, rightIdx-leftIdx+1)
	}
	if maxLen < 3 {
		maxLen = 0
	}
	return maxLen
}

// @lc code=end

