func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 0 {
		return 0
	}
	record1 := make([]int, n+1)
	record2 := make([]int, n+1)
	record1[1] = nums[0]
	for i := 2; i < n+1; i++ {
		if record1[i-1] > record1[i-2]+nums[i-1] {
			record1[i] = record1[i-1]
		} else {
			record1[i] = record1[i-2] + nums[i-1]
		}

		if record2[i-1] > record2[i-2]+nums[i-1] {
			record2[i] = record2[i-1]
		} else {
			record2[i] = record2[i-2] + nums[i-1]
		}
	}
	if record1[n-1] > record2[n] {
		return record1[n-1]
	}
	fmt.Println(record1[n-1], record2[n])
	return record2[n]
}
