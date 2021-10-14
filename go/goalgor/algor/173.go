/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	idx  int
	vals []int
}

func Constructor(root *TreeNode) BSTIterator {
	ans := []int{}
	var iterate func(*TreeNode)
	iterate = func(root *TreeNode) {
		if root == nil {
			return
		}
		iterate(root.Left)
		ans = append(ans, root.Val)
		iterate(root.Right)
	}
	iterate(root)

	return BSTIterator{
		idx:  -1,
		vals: ans,
	}
}

func (this *BSTIterator) Next() int {
	this.idx = this.idx + 1
	return this.vals[this.idx]
}

func (this *BSTIterator) HasNext() bool {
	if this.idx+1 == len(this.vals) {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
