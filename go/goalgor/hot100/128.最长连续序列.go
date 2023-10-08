/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
	maxlen := 1
	curlen := 1
	// startidx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 == nums[i] {
			curlen++
		} else if nums[i-1] == nums[i] {
			continue
		} else {
			maxlen = max(maxlen, curlen)
			curlen = 1
		}
	}

	return max(maxlen, curlen)
}

// @lc code=end

