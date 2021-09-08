func twoSum(nums []int, target int) []int {
	mappings := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mappings[target-nums[i]]; !ok {
			mappings[target-nums[i]] = i
		}
	}
	for i := 0; i < len(nums); i++ {
		if v, ok := mappings[nums[i]]; ok && v != i {
			return []int{v, i}
		}
	}
	return []int{}
}