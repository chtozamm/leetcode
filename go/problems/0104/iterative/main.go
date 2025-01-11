package main

import "container/list"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	level := 0

	for queue.Len() > 0 {
		n := queue.Len()

		for i := 0; i < n; i++ {
			item := queue.Front()
			queue.Remove(item)

			node := item.Value.(*TreeNode)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		level++
	}

	return level
}
