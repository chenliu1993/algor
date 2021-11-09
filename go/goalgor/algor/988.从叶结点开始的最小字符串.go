/*
 * @lc app=leetcode.cn id=988 lang=golang
 *
 * [988] 从叶结点开始的最小字符串
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
func smallestFromLeaf(root *TreeNode) string {
	var (
		str   []string
		child []*TreeNode

		parent    map[*TreeNode]*TreeNode
		getParent func(*TreeNode)

		iterate func(*TreeNode)

		getStr func(*TreeNode, string)
	)
	getParent = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parent[root.Left] = root
			getParent(root.Left)
		}
		if root.Right != nil {
			parent[root.Right] = root
			getParent(root.Right)
		}
	}
	parent = map[*TreeNode]*TreeNode{root: nil}
	getParent(root)

	child = []*TreeNode{}
	iterate = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			child = append(child, root)
			return
		}
		if root.Left != nil {
			iterate(root.Left)
		}
		if root.Right != nil {
			iterate(root.Right)
		}
	}
	iterate(root)

	getStr = func(root *TreeNode, s string) {
		if root == nil {
			str = append(str, s)
			return
		}
		if parent[root] != nil {
			s = s + string('a'+root.Val)
			getStr(parent[root], s)
			s = s[:len(s)-1]
		}
	}
	str = []string{}
	for i := 0; i < len(child); i++ {
		getStr(child[i], "")
	}
	sort.Strings(str)
	return str[0]
}

// @lc code=end

