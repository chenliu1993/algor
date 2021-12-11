/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start

func qsort(nums []int) {
	var sort func(start, end int)
	sort = func(start, end int) {
		if start >= end {
			return
		}
		left := start
		right := end - 1
		standard := nums[end]
		for left < right {
			for left < right && nums[left] <= standard {
				left++
			}
			for left < right && nums[right] >= standard {
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
			sort(start, left-1)
		}
		sort(left+1, end)
	}
	sort(0, len(nums)-1)
}

func triangleNumber(nums []int) int {
	n := len(nums)
	var (
		count int
		cur   int
	)
	count = 0
	// sort.Ints(nums)
	qsort(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cur = nums[i] + nums[j]
			left, right, k := j+1, n-1, j
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

// func triangleNumber(nums []int) (ans int) {
// 	sort.Ints(nums)
// 	for i, v := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			ans += sort.SearchInts(nums[j+1:], v+nums[j])
// 		}
// 	}
// 	return
// }

// func triangleNumber(nums []int) int {
// 	n := len(nums)
// 	var (
// 		count int
// 		cur   int
// 	)
// 	sort.Ints(nums)
// 	count = 0
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			cur = nums[i] + nums[j]
// 			for k := j + 1; k < n; k++ {
// 				if nums[k] < cur {
// 					count++
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

// @lc code=end

