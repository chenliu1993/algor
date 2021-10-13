/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left != nil && root.Right != nil {
			curNode := findMin(root.Right)
			root.Val = curNode.Val
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			curNode := root
			if root.Left == nil {
				root = root.Right
			} else if root.Right == nil {
				root = root.Left
			}
			curNode.Left = nil
			curNode.Right = nil
		}
	}
	return root
}
