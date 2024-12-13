package rotate_image

import "testing"

type testCase struct {
	matrix   [][]int
	expected [][]int
}

func TestRotate(t *testing.T) {
	testCases := []testCase{
		{
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			matrix:   [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			expected: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}

	for _, tc := range testCases {
		rotate(tc.matrix)
		if !matrixEqual(tc.matrix, tc.expected) {
			t.Errorf("got %v, expected %v", tc.matrix, tc.expected)
		}
	}
}

func matrixEqual(a, b [][]int) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
