/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	height := 0
	for i := 0; i < len(root.Children); i++ {
		height = Max(height, maxDepth(root.Children[i]))
	}
	return height + 1
}

// @lc code=end

