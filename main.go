package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(values)
	m := len(queries)
	var (
		start, end, idx int
		adjs            [][]int
		ans             []float64
		route           [][]float64
		dict            map[string]int
		setVal          func(int)
		visited         []int
		sum             float64
	)
	setVal = func(posi int) {
		if posi == end {
			return
		}
		for i := 0; i < len(adjs[posi]); i++ {
			if visited[adjs[posi][i]] == 1 {
				continue
			}
			visited[adjs[posi][i]] = 1
			temp := sum
			sum = sum * route[posi][adjs[posi][i]]
			route[start][adjs[posi][i]] = sum
			setVal(adjs[posi][i])
			sum = temp
		}
	}
	idx = 0
	dict = map[string]int{}
	for i := 0; i < n; i++ {
		if _, ok := dict[equations[i][0]]; !ok {
			dict[equations[i][0]] = idx
			idx++
		}
		if _, ok := dict[equations[i][1]]; !ok {
			dict[equations[i][1]] = idx
			idx++
		}
	}
	route = make([][]float64, idx)
	for i := 0; i < idx; i++ {
		route[i] = make([]float64, idx)
	}
	adjs = make([][]int, idx)
	for i := 0; i < idx; i++ {
		for j := 0; j < idx; j++ {
			route[i][j] = -1.0
		}
		route[i][i] = 1.0
		adjs[i] = []int{}
	}

	for i := 0; i < n; i++ {
		route[dict[equations[i][0]]][dict[equations[i][1]]] = values[i]
		route[dict[equations[i][1]]][dict[equations[i][0]]] = 1 / values[i]
		adjs[dict[equations[i][0]]] = append(adjs[dict[equations[i][0]]], dict[equations[i][1]])
		adjs[dict[equations[i][1]]] = append(adjs[dict[equations[i][1]]], dict[equations[i][0]])
	}
	for i := 0; i < idx; i++ {
		for j := 0; j < idx; j++ {
			if route[i][j] != -1 {
				continue
			}
			start, end = i, j
			visited = make([]int, idx)
			visited[start] = 1
			sum = 1.0
			setVal(start)
		}
	}
	ans = []float64{}
	for i := 0; i < m; i++ {
		if _, ok := dict[queries[i][0]]; !ok {
			ans = append(ans, -1.0)
			continue
		}
		if _, ok := dict[queries[i][1]]; !ok {
			ans = append(ans, -1.0)
			continue
		}
		ans = append(ans, route[dict[queries[i][0]]][dict[queries[i][1]]])
	}
	return ans
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	// equations := [][]string{{"a", "c"}, {"b", "e"}, {"c", "d"}, {"e", "d"}}
	// values := []float64{2.0, 3.0, 0.5, 5.0}
	// queries := [][]string{{"a", "b"}}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(calcEquation(equations, values, queries))
}
