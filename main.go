package main

import (
	"fmt"
)

func maxChunksToSorted(arr []int) int {
	n := len(arr)
	var (
		blocks int
		met    map[int]bool
	)
	blocks = 0
	met = map[int]bool{}
	for i := 0; i < n; i++ {
		// i and arr[i] is grouped
		if _, ok := met[i]; !ok {
			met[i] = true
		} else {
			delete(met, i)
		}
		if _, ok := met[arr[i]]; !ok {
			met[arr[i]] = true
		} else {
			delete(met, arr[i])
		}
		if len(met) == 0 {
			blocks++
		}
	}
	return blocks
}
func main() {
	arr := []int{4, 3, 2, 1, 0}
	fmt.Println(maxChunksToSorted(arr))
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	// fmt.Println(removeComments(source))
}
