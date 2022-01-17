package main

import (
	"container/heap"
	"fmt"
)

type ele struct {
	val   int
	count int
}

type eles []*ele

func (es eles) Less(i, j int) bool {
	return es[i].count > es[j].count
}

func (es eles) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func (es eles) Len() int {
	return len(es)
}

func (es *eles) Push(x interface{}) {
	item := x.(*ele)
	*es = append(*es, item)
}

func (es *eles) Pop() interface{} {
	old := *es
	n := len(old)
	e := old[n-1]
	old[n-1] = nil
	*es = old[:n-1]
	return e
}

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
	if n == 1 {
		return barcodes
	}
	es := eles{}
	record := map[int]int{}
	for _, v := range barcodes {
		if _, ok := record[v]; !ok {
			record[v] = 1
		} else {
			record[v]++
		}
	}
	for k, v := range record {
		heap.Push(&es, &ele{
			val:   k,
			count: v,
		})
	}
	ans := make([]int, n)
	var a *ele
	i := 0
	outbound := false
	for i < n {
		a = heap.Pop(&es).(*ele)
		for a.count != 0 {
			ans[i] = a.val
			i = i + 2
			a.count--
			if a.count == 0 {
				break
			}
			if i >= n {
				outbound = true
				break
			}
		}
		if outbound {
			break
		}
	}

	i = 1
	for i < n {
		if a != nil {
			for a.count != 0 {
				ans[i] = a.val
				i = i + 2
				a.count--
				if a.count == 0 {
					break
				}
				if i >= n {
					outbound = true
					break
				}
			}
		}
		a = heap.Pop(&es).(*ele)
		for a.count != 0 {
			ans[i] = a.val
			i = i + 2
			a.count--
			if a.count == 0 {
				break
			}
			if i >= n {
				outbound = true
				break
			}
		}
		if outbound {
			heap.Push(&es, a)
		}
	}
	return ans
}

func main() {
	barcodes := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	fmt.Println(rearrangeBarcodes(barcodes))
}
