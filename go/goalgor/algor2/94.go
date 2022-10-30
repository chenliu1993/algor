package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := append([]int{}, inorderTraversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal(root.Right)...)
	return ans
}
