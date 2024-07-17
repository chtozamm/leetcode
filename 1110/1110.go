// 1110. Delete Nodes And Return Forest
package forest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	toDeleteMap := make(map[int]bool)
	for _, n := range to_delete {
		toDeleteMap[n] = true
	}

	forest := []*TreeNode{}

	root, forest = processNode(root, toDeleteMap, forest)

	if root != nil {
		forest = append(forest, root)
	}

	return forest
}

func processNode(node *TreeNode, toDeleteMap map[int]bool, forest []*TreeNode) (*TreeNode, []*TreeNode) {
	if node == nil {
		return nil, forest
	}

	node.Left, forest = processNode(node.Left, toDeleteMap, forest)
	node.Right, forest = processNode(node.Right, toDeleteMap, forest)

	if toDeleteMap[node.Val] {
		if node.Left != nil {
			forest = append(forest, node.Left)
		}
		if node.Right != nil {
			forest = append(forest, node.Right)
		}
		return nil, forest
	}

	return node, forest
}
