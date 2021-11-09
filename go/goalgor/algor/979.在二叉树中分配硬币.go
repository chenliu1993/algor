/*
 * @lc app=leetcode.cn id=979 lang=golang
 *
 * [979] 在二叉树中分配硬币
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

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func distributeCoins(root *TreeNode) int {
	var (
		steps      int
		splitCoins func(*TreeNode) int
	)
	splitCoins = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := splitCoins(root.Left)
		right := splitCoins(root.Right)
		steps = steps + Abs(left) + Abs(right)
		return root.Val + left + right - 1
	}
	// Begin split coins
	steps = 0
	splitCoins(root)
	return steps
}

// @lc code=end

