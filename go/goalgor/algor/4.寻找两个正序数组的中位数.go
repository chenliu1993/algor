/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start

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

// @lc code=end

