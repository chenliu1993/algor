/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 */

// @lc code=start

type pair struct {
	content []int
	sum     int
}

type pairs []pair

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (ps pairs) Less(i, j int) bool {
	if ps[i].sum < ps[j].sum {
		return true
	}
	return false
}

func (ps pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps pairs) Len() int {
	return len(ps)
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var (
		m, n  int
		ans   [][]int
		store pairs
	)
	m = Min(len(nums1), k)
	n = Min(len(nums2), k)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			store = append(store, pair{
				content: []int{nums1[i], nums2[j]},
				sum:     nums1[i] + nums2[j],
			})
		}
	}
	sort.Sort(store)
	ans = [][]int{}
	for i := 0; i < Min(k, store.Len()); i++ {
		ans = append(ans, store[i].content)
	}
	return ans
}

// @lc code=end

