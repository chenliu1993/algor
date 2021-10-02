// func largestRectangleArea(heights []int) int {
// 	n := len(heights)
// 	var (
// 		left  int
// 		right int
// 		ans   int
// 		j     int
// 	)
// 	for i := 0; i < n; i++ {
// 		left = i
// 		right = i
// 		for j = i + 1; j < n; j++ {
// 			if heights[j] < heights[i] {
// 				right = j - 1
// 				break
// 			}
// 		}
// 		if j == n {
// 			right = n - 1
// 		}
// 		for j = i - 1; j >= 0; j-- {
// 			if heights[j] < heights[i] {
// 				left = j + 1
// 				break
// 			}
// 		}
// 		if j < 0 {
// 			left = 0
// 		}
// 		if ans < ((right - left + 1) * heights[i]) {
// 			ans = ((right - left + 1) * heights[i])
// 		}
// 	}
// 	return ans
// }

// Brute force n*n
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
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stk := []int{}
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stk) != 0 && heights[stk[len(stk)-1]] >= heights[i] {
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
	for i := n - 1; i >= 0; i-- {
		for len(stk) != 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			right[i] = n
		} else {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = Max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func main() {
	nums := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(nums))
}
