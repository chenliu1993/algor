/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	var iterate func(*TreeNode)
	iterate = func(root *TreeNode) {
		if root.Val >= low && root.Val <= high {
			sum = sum + root.Val
		}
		if root.Left != nil {
			iterate(root.Left)
		}
		if root.Right != nil {
			iterate(root.Right)
		}
	}
	iterate(root)
	return sum
}