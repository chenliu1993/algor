/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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

func pathSum(root *TreeNode, targetSum int) [][]int {
	var dfs func(node *TreeNode, targetSum int)
	var ans [][]int
	var path []int
	dfs = func(node *TreeNode, targetSum int) {
		if node == nil {
			return
		}

		targetSum -= node.Val
		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil && targetSum == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
		}

		dfs(node.Left, targetSum)
		dfs(node.Right, targetSum)

		path = path[:len(path)-1]
	}

	dfs(root, targetSum)

	return ans

}

// @lc code=end

