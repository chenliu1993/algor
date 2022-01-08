package main

import (
	"fmt"
	"sort"
)

// func Less(trip1, trip2 []int) bool {
// 	if trip1[2] == trip2[2] {
// 		return trip1[1] < trip2[1]
// 	}
// 	return trip1[2] < trip2[2]
// }

// func Swap(trip1, trip2 []int) {
// 	temp := make([]int, len(trip1))
// 	copy(temp, trip1)
// 	for i := 0; i < 3; i++ {
// 		trip1[i] = trip2[i]
// 		trip2[i] = temp[i]
// 	}
// }

// func hsort(trips [][]int, start, end int) {
// 	for {
// 		child := 2*start + 1
// 		if child > end {
// 			break
// 		}
// 		if child+1 <= end && Less(trips[child], trips[child+1]) {
// 			child = child + 1
// 		}
// 		if Less(trips[start], trips[child]) {
// 			Swap(trips[start], trips[child])
// 			start = child
// 		} else {
// 			break
// 		}
// 	}
// }

// func Sort(trips [][]int) {
// 	n := len(trips) - 1
// 	for root := n / 2; root >= 0; root-- {
// 		hsort(trips, root, n)
// 	}
// 	for end := n; end >= 0; end-- {
// 		if Less(trips[end], trips[0]) {
// 			Swap(trips[0], trips[end])
// 			hsort(trips, 0, end-1)
// 		}
// 	}
// }

// func Min(i, j int) int {
// 	if i < j {
// 		return i
// 	}
// 	return j
// }

func carPooling(trips [][]int, capacity int) bool {
	n := len(trips)
	sum := make([]int, 1001)
	for i := 0; i < n; i++ {
		for j := trips[i][1]; j < trips[i][2]; j++ {
			sum[j] += trips[i][0]
		}
	}
	sort.Ints(sum)
	if sum[len(sum)-1] > capacity {
		return false
	}
	return true
}

func main() {
	// trips := [][]int{{8, 2, 3}, {4, 1, 3}, {1, 3, 6}, {8, 4, 6}, {4, 4, 8}}
	trips := [][]int{{9, 3, 6}, {8, 1, 7}, {6, 6, 8}, {8, 4, 9}, {4, 2, 9}}
	capacity := 28
	fmt.Println(carPooling(trips, capacity))
}
