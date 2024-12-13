package water

import (
	"testing"
)

type testCase struct {
	height   []int
	expected int
}

func TestMaxArea(t *testing.T) {
	testCases := []testCase{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expected: 49},
		{height: []int{1, 1}, expected: 1},
	}

	for _, tc := range testCases {
		if got := maxArea(tc.height); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
