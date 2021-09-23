func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost {
			if rightmost < i+nums[i] {
				rightmost = i + nums[i]
			}
			if rightmost >= n-1 {
				return true
			}
		}

	}
	return false
}