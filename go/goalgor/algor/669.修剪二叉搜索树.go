/*
 * @lc app=leetcode.cn id=669 lang=golang
 *
 * [669] 修剪二叉搜索树
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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var (
		reorgnize func(*TreeNode) *TreeNode
	)
	reorgnize = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val < low {
			root = reorgnize(root.Right)
		} else if root.Val > high {
			root = reorgnize(root.Left)
		} else {
			root.Left = reorgnize(root.Left)
			root.Right = reorgnize(root.Right)
		}
		return root
	}
	return reorgnize(root)
}

// @lc code=end

