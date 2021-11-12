/*
 * @lc app=leetcode.cn id=971 lang=golang
 *
 * [971] 翻转二叉树以匹配先序遍历
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
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var (
		idx        int
		impossible bool
		ans        []int
		iterate    func(*TreeNode)
	)
	iterate = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val != voyage[idx] {
			impossible = true
			return
		}
		idx++
		if root.Left != nil && root.Left.Val != voyage[idx] {
			ans = append(ans, root.Val)
			iterate(root.Right)
			iterate(root.Left)
		} else {
			iterate(root.Left)
			iterate(root.Right)
		}
	}
	ans = []int{}
	idx = 0
	impossible = false
	iterate(root)
	if impossible {
		ans = []int{-1}
	}
	return ans
}

// @lc code=end

