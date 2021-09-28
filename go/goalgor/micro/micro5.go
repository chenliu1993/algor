package main

import (
	"fmt"
	"math"
	"sort"
)

// Different Balls with weight
// Order, array ball
// In boxes, with weight max.
// boxes number.
// array, boxes k
// minimum weight of boxes

func countMinWeight(balls []int, boxes int) int {
	n := len(balls) // 4

	sort.Slice(balls, func(i, j int) bool {
		if balls[i] < balls[j] {
			return true
		}
		return false
	})

	minBoxWeight := balls[len(balls)-1]
	maxBoxWeight := 0
	for i := 0; i < n; i++ {
		maxBoxWeight = maxBoxWeight + balls[i]
	}
	totalWeight := maxBoxWeight
	ans := maxBoxWeight
	mid := 0
	curBoxes := 0
	var findRightNum func(int, int)
	findRightNum = func(start, end int) {
		if start >= end {
			return
		}
		if curBoxes == boxes && ans > mid {
			ans = mid
		}
		mid = (start + end) / 2
		curBoxes = int(math.Ceil(float64(totalWeight) / float64(mid)))

		findRightNum(start, mid-1)
		findRightNum(mid+1, end)
	}
	findRightNum(minBoxWeight, maxBoxWeight)
	return ans
}

func main() {
	balls := []int{2, 3, 1, 4}
	boxes := 2
	fmt.Println(countMinWeight(balls, boxes))
}
