package binary_tree

import (
	"testing"
)

type testCase struct {
	descriptions [][]int
	expected     *TreeNode
}

func Test(t *testing.T) {
	testCases := []testCase{
		{
			descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
			expected: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 17,
					},
				},
				Right: &TreeNode{
					Val: 80,
					Left: &TreeNode{
						Val: 19,
					},
				},
			},
		},
		{
			descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		result := createBinaryTree(tc.descriptions)
		if !treesAreEqual(result, tc.expected) {
			t.Errorf("got %v, expected %v", result, tc.expected)
		}
	}
}

func treesAreEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && treesAreEqual(a.Left, b.Left) && treesAreEqual(a.Right, b.Right)
}
