package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	xstr := strconv.Itoa(x)
	if xstr[0] == '-' {
		return false
	}
	left := 0
	right := len(xstr) - 1
	for left <= right {
		if xstr[left] == xstr[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}
func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
