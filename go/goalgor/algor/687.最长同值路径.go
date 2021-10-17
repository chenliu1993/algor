/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func longestUnivaluePath(root *TreeNode) int {
	var (
		maxSum     int
		maxSubPath func(*TreeNode) int
	)
	maxSum = 0
	maxSubPath = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxSubPath(root.Left)
		right := maxSubPath(root.Right)
		leftLength := 0
		rightLength := 0
		if root.Left != nil && root.Left.Val == root.Val {
			leftLength = leftLength + left + 1
		}

		if root.Right != nil && root.Right.Val == root.Val {
			rightLength = rightLength + right + 1
		}
		maxSum = Max(maxSum, leftLength+rightLength)
		return Max(leftLength, rightLength)
	}
	maxSubPath(root)
	return maxSum
}

// @lc code=end

