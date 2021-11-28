/*
 * @lc app=leetcode.cn id=341 lang=golang
 *
 * [341] 扁平化嵌套列表迭代器
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	head     int
	length   int
	holdings []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	iterator := &NestedIterator{
		head:     -1,
		length:   0,
		holdings: []int{},
	}

	for _, v := range nestedList {
		if v.IsInteger() {
			iterator.holdings = append(iterator.holdings, v.GetInteger())
			iterator.length++
		} else {
			newIterator := Constructor(v.GetList())
			iterator.holdings = append(iterator.holdings, newIterator.holdings...)
			iterator.length = iterator.length + newIterator.length
		}
	}
	return iterator
}

func (this *NestedIterator) Next() int {
	next := this.head + 1
	this.head = next
	return this.holdings[next]
}

func (this *NestedIterator) HasNext() bool {
	next := this.head + 1
	return next < this.length
}

// @lc code=end

