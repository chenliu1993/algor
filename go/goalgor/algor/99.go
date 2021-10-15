/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	var (
		ans     []int
		iterate func(*TreeNode)
	)
	iterate = func(root *TreeNode) {
		if root == nil {
			return
		}
		iterate(root.Left)
		ans = append(ans, root.Val)
		iterate(root.Right)
	}
	iterate(root)
	x, y := findSwapped(ans)
	recover(root, 2, x, y)
}

func findSwapped(nums []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			y = i + 1
			if x == -1 {
				x = i
			} else {
				break
			}
		}
	}
	return nums[x], nums[y]
}

func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else if root.Val == y {
			root.Val = x
		}
		count = count - 1
		if count == 0 {
			return
		}
	}
	recover(root.Left, 2, x, y)
	recover(root.Right, 2, x, y)
}