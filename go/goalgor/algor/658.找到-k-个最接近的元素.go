/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
func Abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Less(x, y int, standard int) bool {
	xabs := Abs(x, standard)
	yabs := Abs(y, standard)
	if xabs == yabs {
		return x < y
	}
	return xabs < yabs
}

func Sort(nums []int, root, end int, standard int) {
	for {
		child := 2*root + 1
		if child > end {
			break
		}
		if child+1 <= end && Less(nums[child], nums[child+1], standard) {
			child = child + 1
		}
		if Less(nums[root], nums[child], standard) {
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		} else {
			break
		}
	}
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr) - 1
	for i := n / 2; i >= 0; i-- {
		Sort(arr, i, n, x)
	}
	for i := n; i >= 0; i-- {
		if Less(arr[i], arr[0], x) {
			arr[0], arr[i] = arr[i], arr[0]
			Sort(arr, 0, i-1, x)
		}
	}
	sort.Ints(arr[:k])
	return arr[:k]
}

// @lc code=end

