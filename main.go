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
	count = 0
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cur = nums[i] + nums[j]
			left, right := j+1, n-1
			k := j
			for left <= right {
				mid := (left + right) / 2
				if nums[mid] < cur {
					k = mid
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			count = count + k - j
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
