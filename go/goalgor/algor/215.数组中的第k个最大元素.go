/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func heapsort(nums []int, root, end int) {
	for {
		var child = root*2 + 1
		if child > end {
			break
		}
		if child+1 <= end && nums[child] < nums[child+1] {
			child = child + 1
		}
		if nums[root] < nums[child] {
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		} else {
			break
		}
	}
}

func findKthLargest(nums []int, k int) int {
	var n = len(nums) - 1
	for root := n / 2; root >= 0; root-- {
		heapsort(nums, root, n)
	}
	for end := n; end >= 0; end-- {
		if nums[0] > nums[end] {
			nums[0], nums[end] = nums[end], nums[0]
			heapsort(nums, 0, end-1)
		}
	}
	fmt.Println(nums)
	return nums[n+1-k]
}

// @lc code=end

