package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeFromArray(arr []int, i int) *TreeNode {
	var root *TreeNode

	if i < len(arr) && arr[i] != -1 {
		root = &TreeNode{Val: arr[i]}
		root.Left = treeFromArray(arr, i+1)
		root.Right = treeFromArray(arr, i+2)
	}

	return root
}
