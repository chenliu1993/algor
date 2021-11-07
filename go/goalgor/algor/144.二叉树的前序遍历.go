/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	var (
		ans     []int
		iterate func(*TreeNode)
	)

	iterate = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		iterate(root.Left)
		iterate(root.Right)
	}
	ans = []int{}
	iterate(root)
	return ans
}

// @lc code=end

