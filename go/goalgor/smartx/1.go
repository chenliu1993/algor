package main

import "fmt"

// deorder array
// 8 7 6 5 4 3 2 1
// 1 7 2 6 3 5 4

func Comp(nums []int) []int {
	n := len(nums)
	var (
		leftPos, rightPos, rightEnd int
		ans                         []int
	)
	if n%2 == 0 {
		rightEnd = n / 2
	} else {
		rightEnd = n/2 + 1
	}
	ans = []int{}
	leftPos, rightPos = 0, n-1
	for leftPos < rightEnd && rightPos >= rightEnd {
		ans = append(ans, nums[rightPos], nums[leftPos])
		rightPos--
		leftPos++
	}
	if leftPos < rightEnd {
		ans = append(ans, nums[leftPos])
	}
	return ans
}

func main() {
	nums := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(Comp(nums))
}
