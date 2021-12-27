package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)
	hash := func(x int) int {
		return x - 1
	}
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[hash(nums[i])] != nums[i] {
			nums[hash(nums[i])], nums[i] = nums[i], nums[hash(nums[i])]
		}
	}
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	// nums := [][]int{{3, 4, -1, 1}, {1, 2, 0}, {7, 8, 9, 11, 12}}
	// for _, v := range nums {
	// 	fmt.Println(firstMissingPositive(v))
	// }
	nums := []int{3, 4, -1, 1}
	// nums := []int{2147483647, 100000, 1, 3, 2, 4, 5, 6, 7, 100001}
	fmt.Println(firstMissingPositive(nums))

}
