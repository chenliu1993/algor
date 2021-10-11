package main

import "fmt"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func getAllHeights(matrix [][]byte) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		res[0][i] = int(matrix[0][i] - '0')
	}
	for i := 1; i < m; i++ {
		height := make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] = 1 + res[i-1][j]
			}
		}
		copy(res[i], height)
	}
	return res
}

func getMaxArea(nums []int) int {
	m := len(nums)
	stk := []int{}
	left := make([]int, m)
	right := make([]int, m)
	for i := 0; i < m; i++ {
		for len(stk) != 0 && nums[stk[len(stk)-1]] >= nums[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			left[i] = -1
		} else {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	stk = []int{}
	for i := m - 1; i >= 0; i-- {
		for len(stk) != 0 && nums[stk[len(stk)-1]] >= nums[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			right[i] = m
		} else {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans = Max(ans, (right[i]-left[i]-1)*nums[i])
	}
	return ans
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	ans := 0
	heights := getAllHeights(matrix)
	m := len(heights)
	for i := 0; i < m; i++ {
		ans = Max(ans, getMaxArea(heights[i]))
	}
	return ans
}

func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	fmt.Println(maximalRectangle(matrix))
}
