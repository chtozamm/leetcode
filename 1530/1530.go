// 1530. Number of Good Leaf Nodes Pairs

package good_leafs

/*
	Constraints:

		- The number of nodes in the tree is in the range [1, 2^10]
		- 1 <= Node.val <= 100
		- 1 <= distance <= 10
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	ArraySize        = 12
	MaxDistanceIndex = 10
	TotalCountIndex  = 11
)

func countPairs(root *TreeNode, distance int) int {
	return postOrder(root, distance)[TotalCountIndex] // Count of good pairs is defined at index 11
}

// postOrder returns a slice of the counts of leaf nodes for all possible distances from currentNode.
// The length of the slice is set to be 12.
func postOrder(currentNode *TreeNode, distance int) []int {
	if currentNode == nil {
		return make([]int, ArraySize)
	} else if currentNode.Left == nil && currentNode.Right == nil {
		// If it's a leaf node, then return a slice where the count for leaf nodes with distance 0 is set to 1.
		current := make([]int, ArraySize)
		current[0] = 1
		return current
	}

	// Recur for left and right subtrees
	left := postOrder(currentNode.Left, distance)
	right := postOrder(currentNode.Right, distance)

	current := make([]int, ArraySize)

	// Shift left and right subtree counts by 1 distance and combine them
	for i := 0; i < MaxDistanceIndex; i++ {
		current[i+1] += left[i] + right[i]
	}

	// Initialize to total number of good leaf nodes pairs from left and right subtrees.
	current[TotalCountIndex] = left[TotalCountIndex] + right[TotalCountIndex]

	// Iterate through possible leaf node distance pairs
	for d1 := 0; d1 <= distance; d1++ {
		for d2 := 0; d2 <= distance; d2++ {
			if 2+d1+d2 <= distance {
				// If the total path distance is less than the given distance limit,
				// then add to he total number of good pairs
				current[TotalCountIndex] += left[d1] * right[d2]
			}
		}
	}

	return current
}
