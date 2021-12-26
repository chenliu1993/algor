package algor

func MergeSort(src []int) {
	n := len(src) - 1
	if n < 2 {
		return
	}
	msort(src, 0, n)
}

func msort(src []int, left, right int) {
	var mid int
	if left < right {
		mid = (left + right) / 2
		msort(src, left, mid)
		msort(src, mid+1, right)
		merge(src, left, mid+1, right)
	}
}

func merge(src []int, leftPos, rightPos, rightEnd int) {
	var (
		index           int
		leftEnd         int
		numElments      int
		tempSrc         []int
		leftPosRecorder int
	)
	leftPosRecorder = leftPos
	numElments = rightEnd - leftPos + 1
	tempSrc = make([]int, numElments)
	leftEnd = rightPos - 1
	for {
		if leftPos > leftEnd || rightPos > rightEnd {
			break
		}
		if src[leftPos] <= src[rightPos] {
			tempSrc[index] = src[leftPos]
			index++
			leftPos++
		} else {
			tempSrc[index] = src[rightPos]
			index++
			rightPos++
		}
	}
	for {
		if leftPos > leftEnd {
			break
		}
		tempSrc[index] = src[leftPos]
		index++
		leftPos++
	}
	for {
		if rightPos > rightEnd {
			break
		}
		tempSrc[index] = src[rightPos]
		index++
		rightPos++
	}
	for i := 0; i < numElments; i++ {
		src[leftPosRecorder] = tempSrc[i]
		leftPosRecorder++
	}
}

// func divide(nums []int, start, end int) {
// 	mid := (start + end) / 2
// 	divide(nums, start, mid)
// 	divide(nums, mid+1, end)
// 	merge(nums, start, mid+1, end)
// }

// func merge(nums []int, start, rstart, end int) {
// 	n := end - start + 1
// 	holder := make([]int, n)
// 	var (
// 		ptr               int
// 		leftPtr, rightPtr int
// 	)
// 	ptr = 0
// 	leftPtr = start
// 	rightPtr = rstart
// 	for leftPtr < rstart && rightPtr <= end {
// 		if nums[leftPtr] < nums[rightPtr] {
// 			holder[ptr] = nums[leftPtr]
// 			ptr++
// 			leftPtr++
// 		} else {
// 			holder[ptr] = nums[rightPtr]
// 			ptr++
// 			rightPtr++
// 		}
// 	}
// 	for leftPtr < rstart {
// 		holder[ptr] = nums[leftPtr]
// 		ptr++
// 		leftPtr++
// 	}
// 	for rightPtr <= end {
// 		holder[ptr] = nums[rightPtr]
// 		ptr++
// 		rightPtr++
// 	}
// 	for i := 0; i < n; i++ {
// 		nums[start+i] = holder[i]
// 	}
// }

// func MergeSort(nums []int) {
// 	divide(nums, 0, len(nums)-1)
// }
