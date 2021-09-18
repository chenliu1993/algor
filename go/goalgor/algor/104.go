/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxHeight := 0
	curHeight := 0
	var dfs func(*TreeNode, *int)
	dfs = func(root *TreeNode, curHeight *int) {
		if root == nil {
			if maxHeight < *curHeight {
				maxHeight = *curHeight
			}
			return
		}
		*curHeight = *curHeight + 1
		dfs(root.Left, curHeight)
		dfs(root.Right, curHeight)
		*curHeight = *curHeight - 1
	}
	dfs(root, &curHeight)
	return maxHeight
}
