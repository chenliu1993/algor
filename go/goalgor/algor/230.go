/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	arr := []int{}
	// QianXu iterate
	var preOrder func(*TreeNode)
	preOrder = func(root *TreeNode) {
		arr = append(arr, root.Val)
		if root.Left != nil {
			preOrder(root.Left)
		}
		if root.Right != nil {
			preOrder(root.Right)
		}
	}
	preOrder(root)
	// Heap Sort
	var heapSort func(int, int)
	heapSort = func(root, end int) {
		child := 2*root + 1
		for child <= end {
			if child+1 <= end && arr[child] > arr[child+1] {
				child = child + 1
			}
			if arr[root] > arr[child] {
				arr[root], arr[child] = arr[child], arr[root]
				root = child
				child = 2*root + 1
			} else {
				break
			}
		}
	}
	for i := (len(arr) - 1) / 2; i >= 0; i-- {
		heapSort(i, len(arr)-1)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[0] < arr[i] {
			arr[0], arr[i] = arr[i], arr[0]
			heapSort(0, i-1)
		}
	}
	fmt.Println(arr)
	return arr[len(arr)-k]
}