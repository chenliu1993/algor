package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const INT64_MAX = int64(^uint64(0) >> 1)
const INT64_MIN = ^INT64_MAX

func check(root *TreeNode, low, high int64) bool {
	if root == nil {
		return true
	}
	if int64(root.Val) <= low || int64(root.Val) >= high {
		return false
	}
	return check(root.Left, low, int64(root.Val)) && check(root.Right, int64(root.Val), high)
}

func isValidBST(root *TreeNode) bool {
	return check(root, INT64_MIN, INT64_MAX)
}
