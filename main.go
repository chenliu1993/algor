package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
	Preorder []string
	Inorder  []string
}

func Constructor() Codec {
	return Codec{
		Preorder: []string{},
		Inorder:  []string{},
	}
}

func inorder(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := []string{}
	if root.Left != nil {
		ans = append(ans, inorder(root.Left)...)
	}
	ans = append(ans, strconv.Itoa(root.Val))
	if root.Right != nil {
		ans = append(ans, inorder(root.Right)...)
	}
	return ans
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
	var wg sync.WaitGroup

	wg.Add(2)
	go func(in []string) {
		defer wg.Done()
		in = inorder(root)
	}(this.Inorder)
	go func(pre []string) {
		defer wg.Done()
		pre = preorder(root)
	}(this.Preorder)

	wg.Wait()
	return strings.Join(this.Inorder, ",")
}

func findIdx(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func constructTree(preorder, inorder []string) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal, _ := strconv.Atoi(preorder[0])
	rootIdx := findIdx(inorder, preorder[0])
	root := &TreeNode{
		Val: rootVal,
	}
	inorderLeft := inorder[:rootIdx]
	inorderRight := inorder[rootIdx+1:]
	preorderLeft := preorder[1 : len(inorderLeft)+1]
	preorderRight := preorder[len(inorderLeft)+1:]
	root.Left = constructTree(preorderLeft, inorderLeft)
	root.Right = constructTree(preorderRight, inorderRight)
	return root
}

// Deserializes your encoded data to tree.
// Asume data is inorder
func (this *Codec) deserialize(data string) *TreeNode {
	return constructTree(this.Preorder, strings.Split(data, ","))
}

func main() {
	n := 2
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(numTilings(n))
}
