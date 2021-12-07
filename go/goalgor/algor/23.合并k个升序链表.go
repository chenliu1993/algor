/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hsort(head []*ListNode, root, end int) {
	for {
		child := 2*root + 1
		if child > end {
			break
		}
		if child+1 <= end && head[child].Val > head[child+1].Val {
			child = child + 1
		}
		if head[root].Val > head[child].Val {
			head[root], head[child] = head[child], head[root]
			root = child
		} else {
			break
		}
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	var (
		left []int
		ans  *ListNode
	)

	left = []int{}
	for i := 0; i < n; i++ {
		if lists[i] != nil {
			left = append(left, i)
		}
	}
	tempLists := []*ListNode{}
	for i := 0; i < len(left); i++ {
		tempLists = append(tempLists, lists[left[i]])
	}
	lists = tempLists

	n = len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}

	ans = &ListNode{}
	for i := n / 2; i >= 0; i-- {
		hsort(lists, i, n-1)
	}
	head := ans
	for len(lists) != 1 {
		temp := &ListNode{
			Val: lists[0].Val,
		}
		ans.Next = temp
		ans = temp

		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists = lists[1:]
		}
		for i := len(lists) / 2; i >= 0; i-- {
			hsort(lists, i, len(lists)-1)
		}
	}
	for lists[0] != nil {
		temp := &ListNode{
			Val: lists[0].Val,
		}
		ans.Next = temp
		ans = temp
		lists[0] = lists[0].Next
	}

	return head.Next
}

// @lc code=end

