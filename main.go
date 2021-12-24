package main

import (
	"fmt"
)

func buildHeap(nums []int, start, end int) {
	for {
		child := 2*start + 1
		if child > end {
			break
		}
		if child+1 <= end && nums[child] < nums[child+1] {
			child = child + 1
		}
		if nums[start] < nums[child] {
			nums[start], nums[child] = nums[child], nums[start]
			start = child
		} else {
			break
		}
	}
}

func HeapSort(nums []int) []int {
	end := len(nums) - 1
	for start := end / 2; start >= 0; start-- {
		buildHeap(nums, start, end)
	}
	for last := end; last >= 0; last-- {
		if nums[0] > nums[last] {
			nums[0], nums[last] = nums[last], nums[0]
			buildHeap(nums, 0, last-1)
		}
	}
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	left := start
	right := end - 1
	standard := nums[end]
	for left < right {
		for left < right && nums[left] < standard {
			left++
		}
		for left < right && nums[right] > standard {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	if nums[left] >= nums[end] {
		nums[left], nums[end] = nums[end], nums[left]
	} else {
		left++
	}
	if left != 0 {
		quickSort(nums, start, left-1)
	}
	quickSort(nums, left+1, end)
}

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func BubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}

func sortArray(nums []int) []int {
	return BubbleSort(nums)
}

func main() {
	nums := []int{5, 2, 3, 1}
	// nums := []int{0, 1}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(sortArray(nums))
}

// Over time limit
// func findMaxLength(nums []int) int {
// 	n := len(nums)
// 	var (
// 		maxLen int
// 		sums   []int
// 	)
// 	maxLen = 0
// 	sums = make([]int, n)
// 	sums[0] = nums[0]
// 	for i := 1; i < n; i++ {
// 		sums[i] = sums[i-1] + nums[i]
// 		if sums[i]*2 == i+1 {
// 			maxLen = Max(maxLen, i+1)
// 			continue
// 		}
// 		for j := i - 1; j >= 0; j-- {
// 			if (sums[i]-sums[j])*2 == i-j {
// 				maxLen = Max(maxLen, i-j)
// 			}
// 		}
// 	}
// 	return maxLen
// }
