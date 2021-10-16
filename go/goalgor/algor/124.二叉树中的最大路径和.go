/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxPathSum(root *TreeNode) int {
	var maxGrain func(*TreeNode) int
	maxSum := -300000001
	maxGrain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftGrain := Max(maxGrain(root.Left), 0)
		rightGrain := Max(maxGrain(root.Right), 0)
		newGrain := root.Val + leftGrain + rightGrain
		maxSum = Max(maxSum, newGrain)
		return root.Val + Max(leftGrain, rightGrain)
	}
	maxGrain(root)
	return maxSum
}

// @lc code=end

