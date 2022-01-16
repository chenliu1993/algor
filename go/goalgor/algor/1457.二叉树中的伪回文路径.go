/*
 * @lc app=leetcode.cn id=1457 lang=golang
 *
 * [1457] 二叉树中的伪回文路径
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
func copyMap(dst, src map[int]int) {
	for k, v := range src {
		dst[k] = v
	}
}
func pseudoPalindromicPaths(root *TreeNode) int {
	var (
		count   int
		record  map[int]int
		iterate func(*TreeNode)
	)
	count = 0
	record = map[int]int{}
	iterate = func(root *TreeNode) {
		if root == nil {
			return
		}
		if _, ok := record[root.Val]; !ok {
			record[root.Val] = 1
		} else {
			delete(record, root.Val)
		}

		if root.Left == nil && root.Right == nil {
			if len(record) == 0 || len(record) == 1 {
				count++
			}
			return
		}

		if root.Left != nil {
			temp := map[int]int{}
			copyMap(temp, record)
			iterate(root.Left)
			record = map[int]int{}
			copyMap(record, temp)
		}
		if root.Right != nil {
			temp := map[int]int{}
			copyMap(temp, record)
			iterate(root.Right)
			record = map[int]int{}
			copyMap(record, temp)
		}
	}
	iterate(root)
	return count
}

// @lc code=end

