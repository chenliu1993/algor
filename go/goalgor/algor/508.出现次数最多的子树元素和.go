/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] 出现次数最多的子树元素和
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

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		maxCount int
		ans      []int
		dict     map[int]int
		iterate  func(*TreeNode) int
	)
	dict = map[int]int{}
	maxCount = 0
	iterate = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		cur := iterate(root.Left) + root.Val + iterate(root.Right)
		if _, ok := dict[cur]; !ok {
			dict[cur] = 1
		} else {
			dict[cur]++

		}
		if dict[cur] > maxCount {
			maxCount = dict[cur]
		}
		return cur
	}
	iterate(root)
	ans = []int{}
	for key, value := range dict {
		if value == maxCount {
			ans = append(ans, key)
		}
	}
	return ans
}

// @lc code=end

