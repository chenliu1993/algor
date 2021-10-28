/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
func findBottomLeftValue(root *TreeNode) int {
	var (
		front       int
		rear        int
		layerLength int
		queue       []*TreeNode
	)
	queue = []*TreeNode{root}
	front = 0
	rear = front + 1
	layerLength = 1
	for front != rear {
		end := rear
		layerLength = rear - front
		for i := front; i < end; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				rear++
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				rear++
			}
		}
		front = end
	}
	return queue[len(queue)-layerLength].Val
}

// @lc code=end

