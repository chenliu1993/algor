package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maximumSum(arr []int) int {
	n := len(arr)
	var (
		record, sum []int
	)
	record = make([]int, n)
	sum = make([]int, n)
	sum[0] = arr[0]
	record[0] = arr[0]
	for i := 1; i < n; i++ {
		sum[i] = Max(sum[i-1], arr[i])
		record[i] = Max(sum[i-1], record[i-1]+arr[i])
	}
	return record[n-1]
}

func main() {
	arr := []int{1, -2, 0, 3}
	// nums := []int{1, 2, 3}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(maximumSum(arr))
}
