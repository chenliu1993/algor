package algor

import "fmt"

var (
	oddMin  int
	evenMin int
)

func dfs(src, initial int, lists []int, steps *int) int {
	if src < 0 || src > len(lists)-1 {
		return *steps
	}
	if lists[src] != initial {
		*steps = *steps + 1
	}
	return dfs(src-2, initial, lists, steps) + dfs(src+2, initial, lists, steps)
}

func findMinimumSteps(lists []int) int {
	steps := 0
	// First odd
	for i := 1; i < len(lists); i = i + 2 {
		curMin := dfs(i, lists[i], lists, &steps)
		if oddMin > curMin {
			oddMin = curMin
		}
	}
	steps = 0
	// Then even
	for i := 0; i < len(lists); i = i + 2 {
		curMin := dfs(i, lists[i], lists, &steps)
		if evenMin > curMin {
			evenMin = curMin
		}
	}
	return oddMin + evenMin
}

func main() {
	// var (
	// 	n     int
	// 	lists []int
	// )
	// fmt.Scanln("%d", &n)
	// lists = make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scanln("%d", &lists[i])
	// }
	lists := []int{2, 3, 3, 2}
	fmt.Println(findMinimumSteps(lists))
}
