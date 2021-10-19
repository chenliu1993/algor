/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */

// @lc code=start

// 看到题解基本都用的BFS这个方法，但评论区有朋友在询问如何证明，自己做的时候也怀疑过，因此该方法的正确性似乎并不一目了然，自己思考了一下，将自己的思考结果简单讲述一下。
// 证明思路：
// 对当前的图(初始的图或者删掉了前几层叶子节点的图)，我们要找的满足题设的根节点所在位置只有两种可能，一种在当前图的叶子节点(即度为1的节点)，一种为内部节点，若选择某叶子节点1，则该叶子节点1与任意其他叶子节点2的距离必定为叶子1-内部节点x-叶子2，深度为这三段边之和，必大于等于max(内部x-叶子1，内部x-叶子2)，所以相比于叶子节点，解空间只能出现在内部节点，每层情况都是这样，所以我们要剥开叶子节点直到再无可分的内部节点为止。
// 以上

type graph struct {
	childs  []int
	degrees []int
	adj     [][]int
}

func construct(n int, edges [][]int) graph {
	g := graph{}
	g.childs = []int{}
	g.degrees = make([]int, n)
	g.adj = make([][]int, n)
	for i := 0; i < n; i++ {
		g.adj[i] = []int{}
	}
	for i := 0; i < len(edges); i++ {
		g.degrees[edges[i][0]] = g.degrees[edges[i][0]] + 1
		g.degrees[edges[i][1]] = g.degrees[edges[i][1]] + 1
		g.adj[edges[i][0]] = append(g.adj[edges[i][0]], edges[i][1])
		g.adj[edges[i][1]] = append(g.adj[edges[i][1]], edges[i][0])
	}

	for i := 0; i < n; i++ {
		if g.degrees[i] == 1 {
			g.childs = append(g.childs, i)
		}
	}
	return g
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if len(edges) == 0 {
		return []int{0}
	}
	g := construct(n, edges)
	var (
		first int = 0
		rear  int = len(g.childs)
		ans   []int
	)
	for first != rear {
		ans = []int{}
		end := rear
		for i := first; i < end; i++ {
			ans = append(ans, g.childs[i])
			lists := g.adj[g.childs[i]]
			for j := 0; j < len(lists); j++ {
				g.degrees[lists[j]]--
				if g.degrees[lists[j]] == 1 {
					g.childs = append(g.childs, lists[j])
					rear = rear + 1
				}
			}
		}
		first = end
	}
	return ans
}

// @lc code=end

