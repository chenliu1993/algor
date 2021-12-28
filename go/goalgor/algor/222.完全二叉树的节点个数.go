/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	if root.Left != nil {
		count += countNodes(root.Left)
	}
	if root.Right != nil {
		count += countNodes(root.Right)
	}
	return count
}

// @lc code=end

