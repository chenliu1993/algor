package main

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}
	record := make([]int, len(nums)+1)
	record[0] = nums[0]
	record[1] = Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		record[i] = Max(record[i-1], record[i-2]+nums[i])
	}
	return record[len(nums)-1]
}

// func main() {
// 	nums := []int{1, 2, 3, 1}
// 	fmt.Println(rob(nums))
// }
