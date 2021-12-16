func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	if n == 1 {
		return 0
	}
	var (
		max, min int
		sum      int64
	)

	sum = int64(0)
	for i := 0; i < n; i++ {
		max = nums[i]
		min = nums[i]
		for j := i; j < n; j++ {
			max = Max(max, nums[j])
			min = Min(min, nums[j])
			sum = sum + int64(max-min)
		}
	}
	return sum
}