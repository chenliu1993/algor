/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	var (
		cur     string
		ans     []string
		sum     int
		iterate func(*TreeNode)
	)

	iterate = func(root *TreeNode) {
		cur = cur + strconv.Itoa(root.Val)
		tempStr := cur
		if root.Left == nil && root.Right == nil {
			ans = append(ans, cur)
			return
		}
		if root.Left != nil {
			iterate(root.Left)
			cur = tempStr
		}
		if root.Right != nil {
			iterate(root.Right)
			cur = tempStr
		}
	}
	ans = []string{}
	cur = ""
	iterate(root)
	fmt.Println(ans)
	sum = 0
	for i := 0; i < len(ans); i++ {
		sval, err := strconv.Atoi(ans[i])
		if err != nil {
			panic(err)
		}
		sum = sum + sval
	}
	return sum
}

// @lc code=end

