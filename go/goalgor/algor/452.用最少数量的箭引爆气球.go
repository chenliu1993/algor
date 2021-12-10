/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// func Less(i, j []int) bool {
// 	return i[0] < j[0]
// }

// func qsort(arr [][]int, start int, end int) {
// 	if start >= end {
// 		return
// 	}
// 	left := start
// 	right := end - 1
// 	standard := arr[end]
// 	for left < right {
// 		for left < right && Less(arr[left], standard) {
// 			left++
// 		}
// 		for left < right && Less(standard, arr[right]) {
// 			right--
// 		}
// 		arr[left], arr[right] = arr[right], arr[left]
// 	}
// 	if !Less(arr[left], arr[end]) {
// 		arr[left], arr[end] = arr[end], arr[left]
// 	} else {
// 		left++
// 	}
// 	if left != 0 {
// 		qsort(arr, start, left-1)
// 	}
// 	qsort(arr, left+1, end)
// }

// func sort(arr [][]int) {
// 	qsort(arr, 0, len(arr)-1)
// }

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 1 {
		return n
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	// sort(points)
	// fmt.Println(points)
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

// @lc code=end

