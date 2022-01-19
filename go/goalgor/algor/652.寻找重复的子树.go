/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var inOrder func(*TreeNode) string

	ans := []*TreeNode{}
	dict := map[string]int{}
	inOrder = func(root *TreeNode) string {
		if root == nil {
			return "null"
		}
		id := strconv.Itoa(root.Val) + "," + inOrder(root.Left) + "," + inOrder(root.Right)
		if _, ok := dict[id]; !ok {
			dict[id] = 1
		} else {
			dict[id]++
		}
		if dict[id] == 2 {
			ans = append(ans, root)
		}
		return id
	}
	inOrder(root)
	return ans
}

// @lc code=end

