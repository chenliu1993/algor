/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func Less(record map[int]int, x, y int) bool {
	if record[x] < record[y] {
		return true
	}
	return false
}

func heapsort(record map[int]int, nums []int, root, end int) {
	for {
		var child = root*2 + 1
		if child > end {
			break
		}
		if child+1 <= end && Less(record, nums[child], nums[child+1]) {
			child = child + 1
		}
		if Less(record, nums[root], nums[child]) {
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		} else {
			break
		}
	}
}

func topKFrequent(nums []int, k int) []int {
	var (
		n       int
		newNums []int
		record  map[int]int
	)
	n = len(nums) - 1
	record = map[int]int{}
	for i := 0; i < n+1; i++ {
		if _, ok := record[nums[i]]; !ok {
			record[nums[i]] = 1
		} else {
			record[nums[i]]++
		}
	}
	newNums = []int{}
	for k := range record {
		newNums = append(newNums, k)
	}
	n = len(newNums) - 1
	for root := n; root >= 0; root-- {
		heapsort(record, newNums, root, n)
	}
	for end := n; end >= 0; end-- {
		if !Less(record, newNums[0], newNums[end]) {
			newNums[0], newNums[end] = newNums[end], newNums[0]
			heapsort(record, newNums, 0, end-1)
		}
	}
	if k >= n+1 {
		return newNums
	}
	return newNums[n+1-k:]
}

// @lc code=end

