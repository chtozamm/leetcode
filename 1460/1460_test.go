package equal_arrays

import (
	"testing"
)

type testCase struct {
	target   []int
	arr      []int
	expected bool
}

func TestCanBeEqual(t *testing.T) {
	testCases := []testCase{
		{target: []int{1, 2, 3, 4}, arr: []int{2, 4, 1, 3}, expected: true},
		{target: []int{7}, arr: []int{7}, expected: true},
		{target: []int{3, 7, 9}, arr: []int{3, 7, 11}, expected: false},
	}

	for _, tc := range testCases {
		got := canBeEqual(tc.target, tc.arr)
		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
