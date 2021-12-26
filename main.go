package main

import (
	"fmt"
)

// func divide(nums []int, start, end int) {
// 	mid := (start + end) / 2
// 	divide(nums, start, mid)
// 	divide(nums, mid+1, end)
// 	merge(nums, start, mid+1, end)
// }

// func merge(nums []int, start, rstart, end int) {
// 	n := end - start + 1
// 	holder := make([]int, n)
// 	var (
// 		ptr               int
// 		leftPtr, rightPtr int
// 	)
// 	ptr = 0
// 	leftPtr = start
// 	rightPtr = rstart
// 	for leftPtr < rstart && rightPtr <= end {
// 		if nums[leftPtr] < nums[rightPtr] {
// 			holder[ptr] = nums[leftPtr]
// 			ptr++
// 			leftPtr++
// 		} else {
// 			holder[ptr] = nums[rightPtr]
// 			ptr++
// 			rightPtr++
// 		}
// 	}
// 	for leftPtr < rstart {
// 		holder[ptr] = nums[leftPtr]
// 		ptr++
// 		leftPtr++
// 	}
// 	for rightPtr <= end {
// 		holder[ptr] = nums[rightPtr]
// 		ptr++
// 		rightPtr++
// 	}
// 	for i := 0; i < n; i++ {
// 		nums[start+i] = holder[i]
// 	}
// }

// func MergeSort(nums []int) {
// 	if len(nums) < 2 {
// 		return
// 	}
// 	divide(nums, 0, len(nums)-1)
// }

func returnTarget(nums []int, isOdd bool) float64 {
	if isOdd {
		return float64(nums[len(nums)-1])
	}
	return float64((nums[len(nums)-1] + nums[len(nums)-2])) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	holder := []int{}
	var (
		isOdd             bool
		leftPtr, rightPtr int
		spliter           int
	)
	if (m+n)%2 == 0 {
		isOdd = false
	} else {
		isOdd = true
	}
	spliter = (m + n) / 2
	leftPtr = 0
	rightPtr = 0
	for leftPtr < m && rightPtr < n {
		if nums1[leftPtr] < nums2[rightPtr] {
			holder = append(holder, nums1[leftPtr])
			leftPtr++
		} else {
			holder = append(holder, nums2[rightPtr])
			rightPtr++
		}
		if len(holder) == spliter+1 {
			return returnTarget(holder, isOdd)
		}
	}
	for leftPtr < m {
		holder = append(holder, nums1[leftPtr])
		leftPtr++
		if len(holder) == spliter+1 {
			return returnTarget(holder, isOdd)
		}
	}
	for rightPtr < n {
		holder = append(holder, nums2[rightPtr])
		rightPtr++
		if len(holder) == spliter+1 {
			return returnTarget(holder, isOdd)
		}
	}
	// Should not been here
	return float64(0)
}

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
