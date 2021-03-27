package algor

// let us just waste space for now......
func insertEleMin(heap []int, target int) []int {
	n := len(heap)
	newHeap := make([]int, n+1)
	copy(newHeap, heap)
	var i = (n + 1) / 2
	if i < 2 {
		if newHeap[0] > target {
			newHeap[1] = newHeap[0]
			newHeap[0] = target
		} else {
			newHeap[1] = target
		}
		return newHeap
	}
	for ; newHeap[(i/2)-1] > target; i = i / 2 {
		newHeap[i-1] = newHeap[(i/2)-1]
	}
	newHeap[i-1] = target
	return newHeap
}

func buildMinHeap(src []int) []int {
	var (
		formerHeap []int
	)
	formerHeap = make([]int, 1)
	formerHeap[0] = src[0]
	for i := 1; i < len(src); i++ {
		formerHeap = insertEleMin(formerHeap, src[i])
	}
	return formerHeap
}

func deleteMin(heap []int) (int, []int) {
	var (
		minElementIdx int
		minElement    int
		pos           int
		newSrc        []int
	)
	minElement = heap[minElementIdx]
	for i := 1; i < len(heap); i++ {
		if heap[i] <= minElement {
			minElement = heap[i]
			minElementIdx = i
		}
	}
	newSrc = make([]int, len(heap)-1)
	// reconstruct
	for i := 0; i < len(heap); i++ {
		if i == minElementIdx {
			continue
		}
		newSrc[pos] = heap[i]
		pos++
	}
	newSrc = buildMinHeap(newSrc)
	return minElement, newSrc
}

func HeapSort(src []int) []int {
	var (
		curMinElement int
		idx           int
		res           []int
	)
	res = make([]int, len(src))
	heapSrc := buildMinHeap(src)
	for {
		if len(heapSrc) == 0 {
			break
		}
		curMinElement, heapSrc = deleteMin(heapSrc)
		res[idx] = curMinElement
		idx++
	}
	return res
}
