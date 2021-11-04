/*
 * @lc app=leetcode.cn id=1615 lang=golang
 *
 * [1615] 最大网络秩
 */

// @lc code=start
type graph struct {
	adj [][]int
}

func generateG(n int, roads [][]int) graph {
	g := graph{}
	g.adj = make([][]int, n)
	for i := 0; i < n; i++ {
		g.adj[i] = []int{}
	}
	for i := 0; i < len(roads); i++ {
		g.adj[roads[i][0]] = append(g.adj[roads[i][0]], roads[i][1])
		g.adj[roads[i][1]] = append(g.adj[roads[i][1]], roads[i][0])
	}
	return g
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maximalNetworkRank(n int, roads [][]int) int {
	var (
		ans         int
		g           graph
		networkRank func(int, int) int
	)
	g = generateG(n, roads)
	fmt.Println(g.adj)
	networkRank = func(c1, c2 int) int {
		ranks := 0
		// First count c1
		for i := 0; i < len(g.adj[c1]); i++ {
			ranks++
		}
		// Then count c2
		for i := 0; i < len(g.adj[c2]); i++ {
			if g.adj[c2][i] == c1 {
				continue
			}
			ranks++
		}
		return ranks
	}
	ans = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = Max(ans, networkRank(i, j))
		}
	}
	return ans
}

// @lc code=end

