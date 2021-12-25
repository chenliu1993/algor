package main

import (
	"fmt"
)

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	var (
		flag             bool
		stk, left, right []int
	)
	stk, left, right = []int{}, []int{}, []int{}
	for i := 0; i < n; i++ {
		right = append(right, n)
		left = append(left, -1)
	}
	for i := 0; i < n; i++ {
		for len(stk) != 0 && nums[stk[len(stk)-1]] < nums[i] {
			right[stk[len(stk)-1]] = i
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	stk = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stk) != 0 && nums[stk[len(stk)-1]] > nums[i] {
			left[stk[len(stk)-1]] = i
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	for i := 0; i < n; i++ {
		if left[i] != -1 && right[i] != n {
			flag = true
			break
		}
	}
	return flag
}

func main() {
	nums := []int{2, 1, 5, 0, 4, 6}
	// nums := []int{20, 100, 10, 12, 5, 13}
	// nums := []int{1, 2, 1, 3}
	// nums := []int{0, 1}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(increasingTriplet(nums))
}
