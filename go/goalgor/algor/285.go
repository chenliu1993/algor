/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	ans := []*TreeNode{}
	posi := -1
	cur := -1
	var iterate func(*TreeNode)
	iterate = func(root *TreeNode) {
		if root == nil {
			return
		}
		iterate(root.Left)
		cur = cur + 1
		if root.Val == p.Val {
			posi = cur
		}
		ans = append(ans, root)
		iterate(root.Right)
	}
	iterate(root)
	if posi == -1 || posi+1 == len(ans) {
		return nil
	}
	return ans[posi+1]
}
