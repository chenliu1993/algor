/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// To record the parent nodes
	path := map[int]*TreeNode{root.Val: nil}
	var fillInPath func(*TreeNode)
	fillInPath = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			path[root.Left.Val] = root
			fillInPath(root.Left)
		}
		if root.Right != nil {
			path[root.Right.Val] = root
			fillInPath(root.Right)
		}
	}
	fillInPath(root)

	// Find common root
	path[root.Val] = q
	slow := p
	fast := p
	for fast != nil && slow != nil && path[fast.Val] != nil {
		slow = path[slow.Val]
		fast = path[path[fast.Val].Val]
		if slow.Val == fast.Val {
			break
		}
	}

	ptr1 := p
	ptr2 := slow
	for ptr1.Val != ptr2.Val {
		ptr1 = path[ptr1.Val]
		ptr2 = path[ptr2.Val]
	}
	return ptr1
}