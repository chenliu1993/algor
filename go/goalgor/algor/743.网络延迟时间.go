/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 */

// @lc code=start
type graph struct {
	adj     [][]int
	weights [][]int
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func constructG(times [][]int, n int) graph {
	g := graph{}
	g.adj = make([][]int, n+1)
	g.weights = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		g.adj[i] = []int{}
		g.weights[i] = make([]int, n+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			g.weights[i][j] = 101
		}
	}
	for i := 0; i < len(times); i++ {
		g.adj[times[i][0]] = append(g.adj[times[i][0]], times[i][1])
		g.weights[times[i][0]][times[i][1]] = times[i][2]
	}
	return g
}

func networkDelayTime(times [][]int, n int, k int) int {
	g := constructG(times, n)
	g.weights[k][k] = 0
	used := map[int]int{}
	totalTimeUsed := 0
	for i := 1; i < n+1; i++ {
		minialIndex := -1
		for j := 1; j < n+1; j++ {
			if _, ok := used[j]; ok {
				continue
			}
			if minialIndex == -1 || g.weights[k][j] < g.weights[k][minialIndex] {
				minialIndex = j
			}
		}
		used[minialIndex] = 1
		for j := 0; j < len(g.adj[minialIndex]); j++ {
			g.weights[k][g.adj[minialIndex][j]] = Min(g.weights[k][g.adj[minialIndex][j]], g.weights[k][minialIndex]+g.weights[minialIndex][g.adj[minialIndex][j]])
		}
	}
	for i := 1; i <= n; i++ {
		if g.weights[k][i] == 101 {
			return -1
		}
		totalTimeUsed = Max(totalTimeUsed, g.weights[k][i])
	}
	return totalTimeUsed
}

// @lc code=end

