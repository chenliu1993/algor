/*
 * @lc app=leetcode.cn id=385 lang=golang
 *
 * [385] 迷你语法分析器
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	n := len(s)
	var (
		cur  string
		list []*NestedInteger
	)
	cur = ""
	list = []*NestedInteger{}
	for i := 0; i < n; i++ {
		if s[i] == '[' {
			list = append(list, &NestedInteger{})
			cur = ""
		} else if s[i] == ']' {
			if cur != "" {
				curi, _ := strconv.Atoi(cur)
				n := &NestedInteger{}
				n.SetInteger(curi)
				list[len(list)-1].Add(*n)
				cur = ""
			}
			if len(list) > 1 {
				l := list[len(list)-1]
				list = list[:len(list)-1]
				list[len(list)-1].Add(*l)
			}
		} else if s[i] == ',' {
			if cur != "" {
				curi, _ := strconv.Atoi(cur)
				n := &NestedInteger{}
				n.SetInteger(curi)
				list[len(list)-1].Add(*n)
				cur = ""
			}
		} else {
			cur = cur + string(s[i])
		}
	}
	if cur != "" {
		curi, _ := strconv.Atoi(cur)
		n := &NestedInteger{}
		n.SetInteger(curi)
		list = append(list, n)
	}
	return list[len(list)-1]
}

// @lc code=end

