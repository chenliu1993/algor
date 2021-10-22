/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 */

// @lc code=start
type graph struct {
	child   []int
	degrees map[int]int
	adj     map[int][]int
}

func constructG(edges [][]int) graph {
	g := graph{
		child: []int{},
	}
	g.degrees = map[int]int{}
	g.adj = map[int][]int{}
	for i := 0; i < len(edges); i++ {
		g.degrees[edges[i][0]] = g.degrees[edges[i][0]] + 1
		g.degrees[edges[i][1]] = g.degrees[edges[i][1]] + 1
		g.adj[edges[i][0]] = append(g.adj[edges[i][0]], edges[i][1])
		g.adj[edges[i][1]] = append(g.adj[edges[i][1]], edges[i][0])
	}
	for i := 1; i < len(g.degrees)+1; i++ {
		if g.degrees[i] == 1 {
			g.child = append(g.child, i)
		}
	}
	return g
}

func isTree(g graph) bool {
	if len(g.child) == 0 {
		return false
	}
	var (
		front int
		rear  int
	)
	front = 0
	rear = len(g.child)
	for front != rear {
		end := rear
		for i := front; i < end; i++ {
			for j := 0; j < len(g.adj[g.child[i]]); j++ {
				g.degrees[g.adj[g.child[i]][j]] = g.degrees[g.adj[g.child[i]][j]] - 1
				if g.degrees[g.adj[g.child[i]][j]] == 1 {
					g.child = append(g.child, g.adj[g.child[i]][j])
					rear = rear + 1
				}
			}
			front = end
		}
	}
	if len(g.degrees) != len(g.child) {
		// Number of nodes does not meet
		return false
	}
	return true
}

func findRedundantConnection(edges [][]int) []int {
	var (
		i int
		g graph
	)
	for i = len(edges) - 1; i >= 0; i-- {
		tempEdges := make([][]int, len(edges[:i]))
		copy(tempEdges, edges[:i])
		tempEdges = append(tempEdges, edges[i+1:]...)
		g = constructG(tempEdges)
		if isTree(g) {
			break
		}
	}
	return edges[i]
}

// @lc code=end

