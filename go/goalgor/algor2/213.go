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

	dp1 := make([]int, len(nums))
	dp2 := make([]int, len(nums))
	dp1[0] = 0
	dp1[1] = nums[1]

	dp2[0] = nums[0]
	dp2[1] = nums[0]

	for i := 2; i < len(nums); i++ {
		dp1[i] = Max(dp1[i-1], dp1[i-2]+nums[i])
		dp2[i] = Max(dp2[i-1], dp2[i-2]+nums[i])
	}
	return Max(dp1[len(nums)-1], dp2[len(nums)-2])
}

// func main() {
// 	nums := []int{2, 3, 2}
// 	fmt.Println(rob(nums))
// }
