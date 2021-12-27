func findRepeatNumber(nums []int) int {
	n := len(nums)
	hash := func(x int) int {
		return x
	}
	for i := 0; i < n; i++ {
		for nums[i] >= 0 && nums[i] < n && nums[hash(nums[i])] != nums[i] {
			nums[hash(nums[i])], nums[i] = nums[i], nums[hash(nums[i])]
		}
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		if i != nums[i] {
			ans = append(ans, nums[i])
		}
	}

	return ans[0]
}