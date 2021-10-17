/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	var (
		can     bool
		iterate func(*TreeNode, int)
	)
	iterate = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		target = target - root.Val
		if root.Left == nil && root.Right == nil && target == 0 {
			can = true
			return
		}
		iterate(root.Left, target)
		iterate(root.Right, target)
	}
	iterate(root, targetSum)
	return can
}

// @lc code=end

