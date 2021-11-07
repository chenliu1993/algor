/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	var (
		ans     []int
		iterate func(*TreeNode)
	)

	iterate = func(root *TreeNode) {
		if root == nil {
			return
		}

		iterate(root.Left)
		iterate(root.Right)
		ans = append(ans, root.Val)
	}
	ans = []int{}
	iterate(root)
	return ans
}

// @lc code=end

