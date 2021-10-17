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
	var (
		curPath []int
		ans     [][]int
		iterate func(*TreeNode, int)
	)
	iterate = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		curPath = append(curPath, root.Val)
		target = target - root.Val
		fmt.Println(curPath)
		if root.Left == nil && root.Right == nil && target == 0 {
			tempPath := make([]int, len(curPath))
			copy(tempPath, curPath)
			ans = append(ans, tempPath)
		}
		iterate(root.Left, target)
		iterate(root.Right, target)
		curPath = curPath[:len(curPath)-1]
	}
	ans = [][]int{}
	curPath = []int{}
	iterate(root, targetSum)
	return ans
}

// @lc code=end

