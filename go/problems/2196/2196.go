// 2196. Create Binary Tree From Descriptions
package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	childSet := make(map[int]bool)

	getOrCreateNode := func(val int) *TreeNode {
		if nodeMap[val] == nil {
			nodeMap[val] = &TreeNode{Val: val}
		}
		return nodeMap[val]
	}

	for _, d := range descriptions {
		parent, val, isLeft := d[0], d[1], d[2]
		parentNode := getOrCreateNode(parent)
		childNode := getOrCreateNode(val)

		childSet[val] = true

		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
	}

	for val, node := range nodeMap {
		if !childSet[val] {
			return node
		}
	}

	return nil
}
