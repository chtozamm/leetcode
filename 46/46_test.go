package permutations

import (
	"testing"

	"github.com/chtozamm/leetcode/utils"
)

type testCase struct {
	nums     []int
	expected [][]int
}

func TestPermute(t *testing.T) {
	testCases := []testCase{
		{nums: []int{1, 2, 3}, expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{nums: []int{0, 1}, expected: [][]int{{0, 1}, {1, 0}}},
		{nums: []int{1}, expected: [][]int{{1}}},
	}

	for _, tc := range testCases {
		got := permute(tc.nums)
		if !utils.SlicesEqualUnordered2D(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
