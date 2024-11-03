package matrix_zeroes

import "testing"

type testCase struct {
	matrix   [][]int
	expected [][]int
}

func TestSetZeroes(t *testing.T) {
	testCases := []testCase{
		{
			matrix:   [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			matrix:   [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, tc := range testCases {
		setZeroes(tc.matrix)
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
