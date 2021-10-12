/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
func isBalanced(root *TreeNode) bool {
	var getHeight func(*TreeNode) int
	getHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + Max(getHeight(root.Left), getHeight(root.Right))
	}
	if root == nil {
		return true
	}
	if Abs(getHeight(root.Left), getHeight(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}
