/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	var (
		ans        int
		subPathSum func(*TreeNode, int) int
		inOrder    func(*TreeNode)
	)
	subPathSum = func(root *TreeNode, target int) int {
		if root == nil {
			return 0
		}
		count := 0
		if root.Val == target {
			count = count + 1
		}
		count = count + subPathSum(root.Left, target-root.Val)
		count = count + subPathSum(root.Right, target-root.Val)
		return count
	}

	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		ans = ans + subPathSum(root, targetSum)
		inOrder(root.Right)
	}
	ans = 0
	inOrder(root)
	return ans
}

// @lc code=end

