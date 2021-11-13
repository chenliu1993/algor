/*
 * @lc app=leetcode.cn id=1705 lang=golang
 *
 * [1705] 吃苹果的最大数目
 */

// @lc code=start
type bucket struct {
	apples int
	days   int
}

type buckets []bucket

func (bs buckets) Less(i, j int) bool {
	return bs[i].days < bs[j].days
}

func (bs buckets) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func (bs buckets) Len() int {
	return len(bs)
}

func (bs *buckets) Push(b interface{}) {
	*bs = append(*bs, b.(bucket))
}

func (bs *buckets) Pop() interface{} {
	item := (*bs)[len(*bs)-1]
	*bs = (*bs)[:len(*bs)-1]
	return item
}

func eatenApples(apples []int, days []int) int {
	var (
		nums int
		bs   buckets
	)
	n := len(days)
	nums = 0
	bs = buckets{}

	for i := 0; i < n || len(bs) != 0; i++ {
		if i < n && apples[i] != 0 {
			item := bucket{
				days:   days[i] + i,
				apples: apples[i],
			}
			heap.Push(&bs, item)
		}
		for len(bs) != 0 && bs[0].days == i {
			heap.Pop(&bs)
		}
		if len(bs) != 0 && bs[0].apples > 0 {
			nums++
			bs[0].apples--
			if bs[0].apples == 0 {
				heap.Pop(&bs)
			}
		}
	}
	return nums
}

// @lc code=end

