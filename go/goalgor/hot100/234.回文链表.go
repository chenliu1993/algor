/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse(node *ListNode) *ListNode {
	var prev, temp *ListNode
	cur := node
	for cur != nil {
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reverseHalf := reverse(slow)
	for head != nil && reverseHalf != nil {
		if head.Val != reverseHalf.Val {
			return false
		}
		head = head.Next
		reverseHalf = reverseHalf.Next
	}
	return true

}

// @lc code=end

