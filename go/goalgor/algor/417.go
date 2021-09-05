
var direction = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var (
	m int
	n int
)

func toOcean(x, y int, can, heights [][]int) {
	can[x][y] = 1
	for i := 0; i < len(direction); i++ {
		newx := x + direction[i][0]
		newy := y + direction[i][1]
		if newx < 0 || newx > m-1 || newy < 0 || newy > n-1 {
			continue
		}
		if can[newx][newy] == 0 && heights[x][y] <= heights[newx][newy] {
			toOcean(newx, newy, can, heights)
		}
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	m = len(heights)
	n = len(heights[0])
	results := [][]int{}
	canPacific := make([][]int, m)
	canAtlantic := make([][]int, m)
	for i := 0; i < m; i++ {
		canPacific[i] = make([]int, n)
		canAtlantic[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		toOcean(i, 0, canPacific, heights)
	}

	for i := 0; i < n; i++ {
		toOcean(0, i, canPacific, heights)
	}

	for i := 0; i < m; i++ {
		toOcean(i, n-1, canAtlantic, heights)
	}

	for i := 0; i < n; i++ {
		toOcean(m-1, i, canAtlantic, heights)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canPacific[i][j] == 1 && canAtlantic[i][j] == 1 {
				results = append(results, []int{i, j})
			}
		}
	}
	fmt.Println(canPacific)
	fmt.Println("\n")
	fmt.Println(canAtlantic)
	return results
}