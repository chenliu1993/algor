/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
type Intervals [][]int

func (i Intervals) Len() int {
	return len(i)
}

func (i Intervals) Less(x, y int) bool {
	return i[x][0] < i[y][0]
}

func (i Intervals) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func compareIntervals(src1, src2 []int) ([]int, []int) {
	if src2[0] <= src1[1] {
		if src2[1] <= src1[1] {
			return src1, nil
		}
		return []int{src1[0], src2[1]}, nil
	}
	return src1, src2
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Sort(Intervals(intervals))
	res := [][]int{intervals[0]}
	start := 0
	init := 1
	for {
		dst1, dst2 := compareIntervals(res[start], intervals[init])
		if dst2 == nil {
			res = res[:len(res)-1]
			res = append(res, dst1)

		} else {
			res = append(res, dst2)
		}
		start = len(res) - 1
		init++
		if init >= len(intervals) {
			break
		}
	}
	return res
}

// @lc code=end

