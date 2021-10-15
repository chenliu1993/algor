package main

import (
	"fmt"
)

func quicksort(nums []int) {
	n := len(nums)
	var qsort func(int, int)
	qsort = func(left, right int) {
		if left >= right {
			return
		}
		start := left
		end := right - 1
		standard := nums[right]
		for start < end {
			for start < end && nums[start] < standard {
				start = start + 1
			}
			for start < end && nums[end] > standard {
				end = end - 1
			}
			nums[start], nums[end] = nums[end], nums[start]
		}
		if nums[start] >= nums[right] {
			nums[start], nums[right] = nums[right], nums[start]
		} else {
			start = start + 1
		}
		if start != 0 {
			qsort(left, start-1)
		}
		qsort(start+1, right)
	}
	qsort(0, n-1)
}

func main() {
	nums := []int{2, 3, 0, 1, 4}
	quicksort(nums)
	fmt.Println(nums)
}
