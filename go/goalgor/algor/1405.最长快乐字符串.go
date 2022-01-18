/*
 * @lc app=leetcode.cn id=1405 lang=golang
 *
 * [1405] 最长快乐字符串
 */

// @lc code=start
type node struct {
	val   byte
	count int
}

type nodes []*node

func (ns nodes) Less(i, j int) bool {
	return ns[i].count > ns[j].count
}

func (ns nodes) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (ns nodes) Len() int {
	return len(ns)
}

func (ns *nodes) Push(x interface{}) {
	item := x.(*node)
	*ns = append(*ns, item)
}

func (ns *nodes) Pop() interface{} {
	old := *ns
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*ns = old[:n-1]
	return item
}

func longestDiverseString(a int, b int, c int) string {
	ns := nodes{}
	if a != 0 {
		ns = append(ns, &node{
			val:   'a',
			count: a,
		})
	}
	if b != 0 {
		ns = append(ns, &node{
			val:   'b',
			count: b,
		})
	}
	if c != 0 {
		ns = append(ns, &node{
			val:   'c',
			count: c,
		})
	}
	var (
		lastByte     byte
		lastContinue int
		ans          []byte
	)
	heap.Init(&ns)
	ans = []byte{}
	for ns.Len() != 0 {
		first := heap.Pop(&ns).(*node)
		if lastByte == first.val && lastContinue == 2 {
			second := heap.Pop(&ns).(*node)
			ans = append(ans, second.val)
			second.count--
			if second.count != 0 {
				heap.Push(&ns, second)
			}
			lastContinue = 1
			lastByte = second.val
			heap.Push(&ns, first)
		} else if lastByte == first.val && lastContinue != 2 {
			lastContinue++
			lastByte = first.val
			ans = append(ans, first.val)
			first.count--
			if first.count != 0 {
				heap.Push(&ns, first)
			}
		} else {
			lastContinue = 1
			lastByte = first.val
			ans = append(ans, first.val)
			first.count--
			if first.count != 0 {
				heap.Push(&ns, first)
			}
		}
		if ns.Len() == 1 && lastContinue == 2 {
			break
		}
	}
	return string(ans)
}

// @lc code=end

