package pkg

func heapSort(nums []int) {
	end := len(nums) - 1
	hsort := func(l, r int) {
		for {
			c := 2*l + 1
			if c > r {
				break
			}
			if c+1 <= r && nums[c] < nums[c+1] {
				c = c + 1
			}
			if nums[l] < nums[c] {
				nums[l], nums[c] = nums[c], nums[l]
				l = c
			} else {
				break
			}
		}
	}
	for root := end / 2; root >= 0; root-- {
		hsort(root, end)
	}

	for last := end; last >= 0; last-- {
		if nums[last] < nums[0] {
			nums[last], nums[0] = nums[0], nums[last]
			hsort(0, last-1)
		}
	}
}

func quickSort(nums []int) {
	var qsort func(int, int)
	qsort = func(start, end int) {
		if start >= end {
			return
		}
		left := start
		right := end - 1
		standard := nums[end]
		for left < right {
			for left < right && nums[left] < standard {
				left++
			}
			for left < right && standard < nums[right] {
				right--
			}
			nums[left], nums[right] = nums[right], nums[left]
		}
		if nums[end] < nums[left] {
			nums[end], nums[left] = nums[left], nums[end]
		} else {
			left++
		}
		if left != 0 {
			qsort(start, left-1)
		}
		qsort(left+1, end)
	}
	qsort(0, len(nums)-1)
}

func shellSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	for increment := n / 2; increment > 0; increment = increment / 2 {
		for i := increment; i < n; i++ {
			for j := i; j >= increment; j = j - increment {
				if nums[j] < nums[j-increment] {
					nums[j], nums[j-increment] = nums[j-increment], nums[j]
				}
			}
		}
	}
}

func bubbleSort(nums []int) {
	n := len(nums)

	for i := 1; i < n; i++ {
		for j := i; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func insertionSort(nums []int) {
	n := len(nums)

	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		for j := i; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func mergeSort(nums []int) {
	n := len(nums) - 1
	if n < 2 {
		return
	}
	var (
		msort func(int, int)
		merge func(int, int, int)
	)
	merge = func(start, midStart, end int) {
		tempNums := []int{}
		leftPtr := start
		leftEnd := midStart - 1
		rightPtr := midStart
		for leftPtr <= leftEnd && rightPtr <= end {
			if nums[leftPtr] <= nums[rightPtr] {
				tempNums = append(tempNums, nums[leftPtr])
				leftPtr++
			} else {
				tempNums = append(tempNums, nums[rightPtr])
				rightPtr++
			}
		}

		for leftPtr <= leftEnd {
			tempNums = append(tempNums, nums[leftPtr])
			leftPtr++
		}

		for rightPtr <= end {
			tempNums = append(tempNums, nums[rightPtr])
			rightPtr++
		}

		for i := 0; i < len(tempNums); i++ {
			nums[start+i] = tempNums[i]
		}
	}
	msort = func(start, end int) {
		var mid int
		if start < end {
			mid = (start + end) / 2
			msort(start, mid)
			msort(mid+1, end)
			merge(start, mid+1, end)
		}
	}

	msort(0, n)

}

// Sort entrance
// 1: heap
// 2: shell
// 3: quick
// 4: insert
// 5: merge
// default: bubble

func Sort(nums []int, method int) {
	switch method {
	case 1:
		heapSort(nums)
	case 2:
		shellSort(nums)
	case 3:
		quickSort(nums)
	case 4:
		insertionSort(nums)
	case 5:
		mergeSort(nums)
	default:
		bubbleSort(nums)
	}
}
