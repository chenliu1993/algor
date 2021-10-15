package main

import (
	"fmt"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func jump(nums []int) int {
	n := len(nums)
	ans := 10001
	record := make([]int, n)
	for i := 1; i < n; i++ {
		ans = 10001
		for j := i - 1; j >= 0; j-- {
			if i-j <= nums[j] {
				ans = Min(ans, record[j]+1)
			}
		}
		record[i] = ans
	}
	return record[n-1]
}

func main() {
	nums := []int{2, 3, 0, 1, 4}
	fmt.Println(jump(nums))
}
