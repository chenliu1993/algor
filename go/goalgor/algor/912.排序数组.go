/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func buildHeap(nums []int, start, end int) {
	for {
		child := 2*start + 1
		if child > end {
			break
		}
		if child+1 <= end && nums[child] < nums[child+1] {
			child = child + 1
		}
		if nums[start] < nums[child] {
			nums[start], nums[child] = nums[child], nums[start]
			start = child
		} else {
			break
		}
	}
}

func HeapSort(nums []int) []int {
	end := len(nums) - 1
	for start := end / 2; start >= 0; start-- {
		buildHeap(nums, start, end)
	}
	for last := end; last >= 0; last-- {
		if nums[0] > nums[last] {
			nums[0], nums[last] = nums[last], nums[0]
			buildHeap(nums, 0, last-1)
		}
	}
	return nums
}
func sortArray(nums []int) []int {
	return HeapSort(nums)
}

// @lc code=end

