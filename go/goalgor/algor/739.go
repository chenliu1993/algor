package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stk := []int{}
	for i := 0; i < n; i++ {
		for len(stk) != 0 && temperatures[stk[len(stk)-1]] < temperatures[i] {
			ans[stk[len(stk)-1]] = i - stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return ans
}

func main() {
	nums := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(nums))
}
