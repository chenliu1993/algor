package main

import (
	"fmt"
)

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	even := make([]int64, n)
	odd := make([]int64, n)
	odd[0] = int64(nums[0])
	for i := 1; i < n; i++ {
		even[i] = Max(even[i-1], odd[i-1]-int64(nums[i]))
		odd[i] = Max(odd[i-1], even[i-1]+int64(nums[i]))
	}
	fmt.Println(even)
	fmt.Println(odd)
	return Max(even[n-1], odd[n-1])
}

func main() {
	nums := []int{5, 6, 7, 8}
	// nums := []int{1, 2, 3}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(maxAlternatingSum(nums))
}
