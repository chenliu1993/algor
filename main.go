package main

import (
	"fmt"
)

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Less(i, j []int) bool {
	if i[0] == j[0] {
		return i[1] < j[1]
	}
	return i[0] < j[0]
}

func qsort(arr [][]int, start int, end int) {
	if start >= end {
		return
	}
	left := start
	right := end - 1
	standard := arr[end]
	for left < right {
		for left < right && Less(arr[left], standard) {
			left++
		}
		for left < right && Less(standard, arr[right]) {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
		fmt.Println(arr[left], arr[right])
	}
	if !Less(arr[left], arr[end]) {
		arr[left], arr[end] = arr[end], arr[left]
	} else {
		left++
	}
	if left != 0 {
		qsort(arr, start, left-1)
	}
	qsort(arr, left+1, end)
}

func sort(arr [][]int) {
	qsort(arr, 0, len(arr)-1)
}

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 1 {
		return n
	}
	// sort.Slice(points, func(i, j int) bool {
	// 	return points[i][0] < points[j][0]
	// })
	sort(points)
	fmt.Println(points)
	var (
		count int
		// start int
		end int
	)
	// start = points[0][0]
	end = points[0][1]
	count = 1
	for i := 1; i < n; i++ {
		if points[i][0] > end {
			count++
			// start = points[i][0]
			end = points[i][1]
			continue
		}
		// start = points[i][0]
		end = Min(points[i][1], end)
	}
	return count
}
func main() {
	points := [][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(findMinArrowShots(points))
}
