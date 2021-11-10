/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第 K 小的元素
 */

// @lc code=start
var directions = [][]int{{1, 0}, {0, 1}}

func kthSmallest(matrix [][]int, k int) int {
	var (
		n           int
		front, rear int
		store       []int
	)
	n = len(matrix)
	store = []int{matrix[0][0]}
	front = 0
	rear = front + 1
	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}
	queue := [][]int{{0, 0}}
	for front < rear {
		end := rear
		for i := front; i < end; i++ {
			x := queue[i][0]
			y := queue[i][1]
			for j := 0; j < len(directions); j++ {
				newx := x + directions[j][0]
				newy := y + directions[j][1]
				if newx < 0 || newx >= n || newy < 0 || newy >= n || visited[newx][newy] == 1 {
					continue
				}
				visited[newx][newy] = 1
				queue = append(queue, []int{newx, newy})
				store = append(store, matrix[newx][newy])
				rear++
			}
		}
		front = end
	}
	sort.Ints(store)
	fmt.Println(store)
	return store[k-1]
}

// @lc code=end

