/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		first int
		rear  int
		q     []*TreeNode
		cur   []int
		ans   [][]int
	)
	q = []*TreeNode{root}
	first = 0
	rear = first + 1

	for first != rear {
		cur = []int{}
		for i := first; i < rear; i++ {
			cur = append(cur, q[i].Val)
		}
		ans = append(ans, cur)

		end := rear
		for i := first; i < end; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
				rear = rear + 1
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
				rear = rear + 1
			}
		}
		first = end
	}

	return ans
}

// @lc code=end

