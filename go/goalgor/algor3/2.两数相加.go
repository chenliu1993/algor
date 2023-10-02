/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		sum   = 0
		carry = 0
	)
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	rear := head
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + carry
		carry = int(sum / 10)
		node := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		rear.Next = node
		rear = node
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil {
		for ; l1 != nil; l1 = l1.Next {
			sum = l1.Val + carry
			carry = int(sum / 10)
			node := &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			rear.Next = node
			rear = node
		}
	}
	if l2 != nil {
		for ; l2 != nil; l2 = l2.Next {
			sum = l2.Val + carry
			carry = int(sum / 10)
			node := &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			rear.Next = node
			rear = node
		}
	}
	if carry != 0 {
		node := &ListNode{
			Val:  carry,
			Next: nil,
		}
		rear.Next = node
		rear = node
	}
	return head.Next
}

// @lc code=end

