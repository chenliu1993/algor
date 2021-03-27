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
