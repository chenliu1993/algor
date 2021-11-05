/*
 * @lc app=leetcode.cn id=1448 lang=golang
 *
 * [1448] 统计二叉树中好节点的数目
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
func goodNodes(root *TreeNode) int {
	var (
		maxNum  int
		count   int
		iterate func(*TreeNode)
	)
	iterate = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if root.Val >= maxNum {
				count++
			}
			return
		}
		if root.Val >= maxNum {
			count++
		}
		if root.Left != nil {
			tempMax := maxNum
			if maxNum < root.Val {
				maxNum = root.Val
			}
			iterate(root.Left)
			maxNum = tempMax
		}
		if root.Right != nil {
			tempMax := maxNum
			if maxNum < root.Val {
				maxNum = root.Val
			}
			iterate(root.Right)
			maxNum = tempMax
		}
	}
	count = 0
	maxNum = -10001
	iterate(root)
	return count
}

// @lc code=end

