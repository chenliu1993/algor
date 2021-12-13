/*
 * @lc app=leetcode.cn id=449 lang=golang
 *
 * [449] 序列化和反序列化二叉搜索树
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

type Codec struct {
	Ele []string
}

func Constructor() Codec {
	return Codec{
		Ele: []string{},
	}
}

func preorder(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := []string{}
	ans = append(ans, strconv.Itoa(root.Val))
	if root.Left != nil {
		ans = append(ans, preorder(root.Left)...)
	}
	if root.Right != nil {
		ans = append(ans, preorder(root.Right)...)
	}
	return ans
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.Ele = preorder(root)
	return strings.Join(this.Ele, ",")
}

func constructTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &TreeNode{
			Val: arr[0],
		}
	}
	leftIdx, rightIdx := -1, -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[0] {
			rightIdx = i
			break
		}
	}
	if rightIdx == -1 {
		rightIdx = len(arr)
	}
	leftIdx = 1
	root := &TreeNode{
		Val: arr[0],
	}
	root.Left = constructTree(arr[leftIdx:rightIdx])
	root.Right = constructTree(arr[rightIdx:])
	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	res := strings.Split(data, ",")
	ans := []int{}
	for i := 0; i < len(res); i++ {
		t, _ := strconv.Atoi(res[i])
		ans = append(ans, t)
	}
	return constructTree(ans)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// @lc code=end

