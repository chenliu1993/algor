
// Over time limit
func construct(n int, edges [][]int) [][]int {
	linked := make([][]int, n)
	for i := 0; i < n; i++ {
		linked[i] = make([]int, n)
	}
	for i := 0; i < n-1; i++ {
		linked[edges[i][0]][edges[i][1]] = 1
		linked[edges[i][1]][edges[i][0]] = 1
	}
	return linked
}

func findMinHeightTrees(n int, edges [][]int) []int {
	var (
		curHeight       int
		visited         []int
		height          int
		heights         []int
		globalMinHeight int = 20001
		getCurHeight    func(int)
	)
	adj := construct(n, edges)
	getCurHeight = func(root int) {
		p := adj[root]
		for i := 0; i < len(p); i++ {
			if p[i] == 1 && i != root && visited[i] == 0 {
				curHeight = curHeight + 1
				visited[i] = 1
				getCurHeight(i)
				visited[i] = 0
				curHeight = curHeight - 1
			} else {
				if height < curHeight {
					height = curHeight
				}
			}
		}
	}
	heights = make([]int, n)
	for i := 0; i < n; i++ {
		curHeight = 0
		height = 0
		visited = make([]int, n)
		visited[i] = 1
		getCurHeight(i)
		heights[i] = height
		if height < globalMinHeight {
			globalMinHeight = height
		}
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		if heights[i] == globalMinHeight {
			ans = append(ans, i)
		}
	}
	return ans
}
