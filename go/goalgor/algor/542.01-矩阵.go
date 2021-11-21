/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
var directions [][]int = [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	var (
		queue   [][]int
		visited [][]int
		dist    [][]int
	)
	visited = make([][]int, m)
	dist = make([][]int, m)
	queue = [][]int{}
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				visited[i][j] = 1
				queue = append(queue, []int{i, j})
			}
		}
	}

	front := 0
	rear := len(queue)
	for front != rear {
		end := rear
		for i := front; i < end; i++ {
			x := queue[i][0]
			y := queue[i][1]
			for j := 0; j < len(directions); j++ {
				newx := x + directions[j][0]
				newy := y + directions[j][1]
				if newx < 0 || newx >= m || newy < 0 || newy >= n || visited[newx][newy] == 1 {
					continue
				}
				dist[newx][newy] = dist[x][y] + 1
				queue = append(queue, []int{newx, newy})
				visited[newx][newy] = 1
				rear++

			}
		}
		front = end
	}
	return dist
}

// @lc code=end

