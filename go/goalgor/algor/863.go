/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	lists   []int
	treeMap map[int]*TreeNode
)

func toTreeMap(root *TreeNode) {
	if root.Left != nil {
		treeMap[root.Left.Val] = root
		toTreeMap(root.Left)
	}
	if root.Right != nil {
		treeMap[root.Right.Val] = root
		toTreeMap(root.Right)
	}
}

func iterateAsRoot(target, from *TreeNode, k, depth int) {
	if target == nil {
		return
	}
	if depth == k {
		lists = append(lists, target.Val)
		return
	}
	if target.Left != from {
		iterateAsRoot(target.Left, target, k, depth+1)
	}
	if target.Right != from {
		iterateAsRoot(target.Right, target, k, depth+1)
	}
	if treeMap[target.Val] != from {
		iterateAsRoot(treeMap[target.Val], target, k, depth+1)
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	treeMap = map[int]*TreeNode{}
	lists = []int{}

	toTreeMap(root)
	iterateAsRoot(target, nil, k, 0)

	return lists
}
