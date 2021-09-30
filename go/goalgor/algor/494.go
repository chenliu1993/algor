func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	count := 0
	sum := 0
	var iterate func(int)
	iterate = func(start int) {
		if start == n {
			if sum == target {
				count = count + 1
			}
			return
		}
		// plus
		sum = sum + nums[start]
		iterate(start + 1)
		// minus
		sum = sum - 2*nums[start]
		iterate(start + 1)
		sum = sum + nums[start]
	}
	iterate(0)
	return count
}
