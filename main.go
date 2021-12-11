package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	n := len(nums)
	var (
		count int
		cur   int
	)
	sort.Ints(nums)
	count = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cur = nums[i] + nums[j]
			for k := j + 1; k < n; k++ {
				if nums[k] < cur {
					count++
				}
			}
		}
	}
	return count
}
func main() {
	nums := []int{2, 2, 3, 4}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(triangleNumber(nums))
}
