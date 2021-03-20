package main

import "fmt"

var (
	mid   int
	left  int
	right int
)

func hsort(src []int, tgt, left, right int) bool {
	if left > right {
		return false
	} else if left == right {
		if tgt == src[left] {
			return true
		}
		return false
	}
	mid = (left + right) / 2
	if tgt == src[mid] {
		return true
	} else if tgt < src[mid] {
		return hsort(src, tgt, left, mid)
	} else {
		return hsort(src, tgt, mid+1, right)
	}
}

func halfsort(src []int, tgt int) bool {
	return hsort(src, tgt, 0, len(src)-1)
}

func main() {
	target := 8
	src := []int{1, 2, 5, 6, 7, 12, 23}
	fmt.Println(halfsort(src, target))
}
