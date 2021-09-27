package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	left := make([]int, n)
	right := make([]int, n)
	results := []string{}
	ans := []string{}
	var iterate func()
	iterate = func() {
		if len(left) == 0 && len(right) == 0 {
			results = append(results, strings.Join(ans, ""))
			return
		}
		if len(left) != 0 && len(right) == 0 {
			return
		}
		if len(left) < len(right) {
			// Add left
			if len(left) != 0 {
				ans = append(ans, "(")
				left = left[:len(left)-1]
				iterate()
				ans = ans[:len(ans)-1]
				left = append(left, 0)
			}
			// Add Right
			ans = append(ans, ")")
			right = right[:len(right)-1]
			iterate()
			ans = ans[:len(ans)-1]
			right = append(right, 0)
		} else if len(left) == len(right) {
			// Add left
			ans = append(ans, "(")
			left = left[:len(left)-1]
			iterate()
			ans = ans[:len(ans)-1]
			left = append(left, 0)
		}
	}
	iterate()
	return results
}

func main() {
	fmt.Println(generateParenthesis(3))
}
