/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		res         []int
		ans         []*TreeNode
		first, rear int
	)
	first = 0
	rear = first + 1
	ans = []*TreeNode{root}
	res = []int{}
	for first < rear {
		end := rear
		layerMax := ans[first].Val
		for i := first; i < end; i++ {
			layerMax = Max(layerMax, ans[i].Val)
			if ans[i].Left != nil {
				ans = append(ans, ans[i].Left)
				rear++
			}
			if ans[i].Right != nil {
				ans = append(ans, ans[i].Right)
				rear++
			}
		}
		res = append(res, layerMax)
		first = end
	}
	return res
}

// @lc code=end

