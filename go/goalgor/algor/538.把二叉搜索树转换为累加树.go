/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	var (
		record              []int
		dict                map[int]int
		InOrder, SetInOrder func(*TreeNode)
	)

	InOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			InOrder(root.Left)
		}
		record = append([]int{root.Val}, record...)
		if root.Right != nil {
			InOrder(root.Right)
		}
	}
	SetInOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			SetInOrder(root.Left)
		}
		root.Val = dict[root.Val]
		if root.Right != nil {
			SetInOrder(root.Right)
		}
	}
	record = []int{}
	dict = map[int]int{}
	InOrder(root)
	sum := 0
	for i := 0; i < len(record); i++ {
		sum += record[i]
		dict[record[i]] = sum
	}
	SetInOrder(root)
	return root
}

// @lc code=end

