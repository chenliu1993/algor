package main

import (
	"container/heap"
	"fmt"
)

type minus []int

func (m minus) Len() int {
	return len(m)
}
func (m minus) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m minus) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minus) Push(target interface{}) {
	targetVal, _ := target.(int)
	*m = append(*m, targetVal)
}
func (m *minus) Pop() interface{} {
	target := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return target
}

func magicTower(nums []int) int {
	n := len(nums)
	var (
		count  int
		holder *minus
		total  int64 = int64(1)
	)
	for i := 0; i < n; i++ {
		total = total + int64(nums[i])
	}
	if total+int64(1) <= 0 {
		return -1
	}
	holder = &minus{}
	total = int64(1)
	count = 0
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			heap.Push(holder, nums[i])
		}
		total = total + int64(nums[i])
		if total > int64(0) {
			continue
		}
		v, _ := heap.Pop(holder).(int)
		total = total - int64(v)
		count++
	}
	return count
}
func main() {
	arr := []int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}
	// arr := []int{-200, -300, 400, 0}
	// arr := []int{5, -4, 1, -2, -2, -2, 4}
	fmt.Println(magicTower(arr))
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
}
