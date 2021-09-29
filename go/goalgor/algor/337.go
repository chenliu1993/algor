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
func rob(root *TreeNode) int {
	var dfs func(*TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		yes := root.Val + l[1] + r[1]
		no := Max(l[0], l[1]) + Max(r[0], r[1])
		return []int{yes, no}
	}
	ans := dfs(root)
	return Max(ans[0], ans[1])
}