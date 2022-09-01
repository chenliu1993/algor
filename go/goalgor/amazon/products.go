package main

import "fmt"

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func maxSubArray(nums []int32) int64 {
	n := len(nums)
	record := make([][]int64, n)
	for i := 0; i < n; i++ {
		record[i] = make([]int64, n)
		record[i][i] = int64(nums[i])
	}

	var (
		left int64
		cur  int64
		ans  int64
	)
	for j := 0; j < n; j++ {
		left = int64(nums[j])
		for i := j; i >= 1; i-- {
			if int64(nums[i-1]) >= left-1 {
				cur = record[i][j] + left - 1
				left = left - 1
			} else {
				cur = record[i][j] + int64(nums[i-1])
				left = int64(nums[i-1])
			}
			record[i-1][j] = Max(record[i][j], cur)
		}
		ans = Max(ans, record[0][j])
	}
	return ans
}

func main() {
	input := []int32{7, 3, 4, 5, 3}
	fmt.Println((maxSubArray(input)))
}
