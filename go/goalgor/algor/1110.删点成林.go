/*
 * @lc app=leetcode.cn id=1110 lang=golang
 *
 * [1110] 删点成林
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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var (
		ans     []*TreeNode
		dict    map[int]int
		iterate func(*TreeNode) *TreeNode
	)
	iterate = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Left == nil && root.Right == nil {
			if _, ok := dict[root.Val]; !ok {
				return root
			}
			return nil
		}
		if _, ok := dict[root.Val]; ok {
			if iterate(root.Left) != nil {
				ans = append(ans, iterate(root.Left))
			}
			if iterate(root.Right) != nil {
				ans = append(ans, iterate(root.Right))
			}
			return nil
		}
		root.Left = iterate(root.Left)
		root.Right = iterate(root.Right)
		return root
	}
	dict = map[int]int{}
	for i := 0; i < len(to_delete); i++ {
		dict[to_delete[i]] = 1
	}
	if _, ok := dict[root.Val]; ok {
		ans = []*TreeNode{}
	} else {
		ans = []*TreeNode{root}
	}

	iterate(root)
	return ans
}

// @lc code=end

