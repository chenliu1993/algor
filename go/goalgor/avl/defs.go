package main

type AvlNode struct {
	Val    int
	Height int
	Left   *AvlNode
	Right  *AvlNode
}

func height(root *AvlNode) int {
	if root == nil {
		return -1
	}
	return root.Height
}

func insert(root *AvlNode, val int) *AvlNode {
	if root == nil {
		return &AvlNode{
			Val:    val,
			Height: 0,
			Left:   nil,
			Right:  nil,
		}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
		if height(root.Left)-height(root.Right) == 2 {
			if val < root.Left.Val {
				root = SingleRotateWithLeft(root)
			} else {
				root = DoubleRotateWithLeft(root)
			}
		}
	} else if val > root.Val {
		root.Right = insert(root.Right, val)
		if height(root.Left)-height(root.Right) == 2 {
			if val < root.Left.Val {
				root = DoubleRotateWithRight(root)
			} else {
				root = SingleRotateWithRight(root)
			}
		}
	}
	// No update
	root.Height = Max(height(root.Left), height(root.Right))
	return root
}

func SingleRotateWithLeft(root *AvlNode) *AvlNode {
	var node *AvlNode
	node = root.Left
	root.Left = node.Right
	node.Right = root
	root.Height = Max(height(root.Left), height(root.Right)) + 1
	node.Height = Max(height(node.Left), root.Height) + 1
	return node
}

func SingleRotateWithRight(root *AvlNode) *AvlNode {
	var node *AvlNode
	node = root.Right
	root.Right = node.Left
	node.Left = root
	root.Height = Max(height(root.Left), height(root.Right)) + 1
	node.Height = Max(height(node.Right), root.Height) + 1
	return node
}

func DoubleRotateWithLeft(root *AvlNode) *AvlNode {
	root.Left = SingleRotateWithRight(root.Left)
	return SingleRotateWithLeft(root)
}

func DoubleRotateWithRight(root *AvlNode) *AvlNode {
	root.Left = SingleRotateWithLeft(root.Left)
	return SingleRotateWithRight(root)
}
