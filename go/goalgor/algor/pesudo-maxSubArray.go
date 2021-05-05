package algor

import "math"

// Dividing the problem.
// T(n)=2*T(n/2)+O(n)+O(1) since the if-else is seens as a contant time-costing step
// and the findCrossMaxArray only iterates through the src.
// Thus the solution to above T(n) is:
// T(n)=O(n*lg(n)

// T(N) = aT(N/b)+O(N^k)
// T(N) = O(N^(logb(a))) if a>b^k
//      = O((N^k)*lg(N)) if a=b^k
// 	 = O(N^k)         if a<b^k

func findMaxArray(src []int, left, right int) (int, int, int) {
	if left == right {
		return left, right, src[left]
	}
	mid := (left + right) / 2
	leftLowMax, leftHighMax, leftMax := findMaxArray(src, left, mid)
	rightLowMax, rightHighMax, rightMax := findMaxArray(src, mid+1, right)
	crossLowMax, crossHighMax, crossMax := findCrossMaxArray(src, left, mid, right)
	if leftMax >= rightMax && leftMax >= crossMax {
		return leftLowMax, leftHighMax, leftMax
	} else if rightMax >= leftMax && rightMax >= crossMax {
		return rightLowMax, rightHighMax, rightMax
	}
	return crossLowMax, crossHighMax, crossMax
}

func findCrossMaxArray(src []int, left, mid, right int) (int, int, int) {
	var (
		leftCrossMax  int
		rightCrossMax int
	)
	leftMax := -1 * math.MaxInt64
	sum := 0
	for i := mid; i >= left; i-- {
		sum := sum + src[i]
		if sum > leftMax {
			leftMax = sum
			leftCrossMax = i
			break
		}
	}
	rightMax := -1 * math.MaxInt64
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum := sum + src[i]
		if sum > rightMax {
			rightMax = sum
			rightCrossMax = i
			break
		}
	}
	return leftCrossMax, rightCrossMax, (leftMax + rightMax)
}
