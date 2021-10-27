package main

import (
	"fmt"
	"log"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxRotateFunction(nums []int) int {
	n := len(nums)
	var (
		record []int
	)
	record = make([]int, n)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum = sum + j*nums[(j+i)%n]
		}
		if i == 0 {
			record[i] = sum
		} else {
			record[i] = Max(record[i-1], sum)
		}
	}
	return record[n-1]
}

func main() {
	nums := []int{4, 3, 2, 6, 6, 6, 6, 6, 6, 8}
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	fmt.Println(maxRotateFunction(nums))
}

// Over time limit
