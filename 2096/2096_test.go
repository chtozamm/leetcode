package directions

import (
	"testing"
)

type testCase struct {
	root       *TreeNode
	startValue int
	destValue  int
	expected   string
}

func TestGetDirections(t *testing.T) {
	testCases := []testCase{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			startValue: 3,
			destValue:  6,
			expected:   "UURL",
		},
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			startValue: 2,
			destValue:  1,
			expected:   "L",
		},
	}

	for _, tc := range testCases {
		got := getDirections(tc.root, tc.startValue, tc.destValue)
		if got != tc.expected {
			t.Errorf("got %q, expected %q", got, tc.expected)
		}
	}
}
