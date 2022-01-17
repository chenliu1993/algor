/*
 * @lc app=leetcode.cn id=1054 lang=golang
 *
 * [1054] 距离相等的条形码
 */

// @lc code=start
type ele struct {
	val   int
	count int
}

func rearrangeBarcodes(barcodes []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(barcodes); i++ {
		m[barcodes[i]]++
	}
	n := make([]ele, 0, len(m))
	for k, v := range m {
		n = append(n, ele{k, v})
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i].count > n[j].count
	})
	ret := make([]int, len(barcodes))
	start := 0
	for _, v := range n {
		val, count := v.val, v.count
		for count > 0 {
			ret[start] = val
			start += 2
			count--
			if start >= len(barcodes) {
				start = 1
			}
		}
	}
	return ret
}

// @lc code=end

