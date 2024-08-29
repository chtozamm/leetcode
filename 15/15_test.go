package three_sum

import (
	"testing"

	"github.com/chtozamm/leetcode/utils"
)

type testCase struct {
	nums     []int
	expected [][]int
}

func TestThreeSum(t *testing.T) {
	testCases := []testCase{
		{nums: []int{-1, 0, 1, 2, -1, -4}, expected: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{nums: []int{0, 1, 1}, expected: [][]int{}},
		{nums: []int{0, 0, 0}, expected: [][]int{{0, 0, 0}}},
	}

	for _, tc := range testCases {
		if got := threeSum(tc.nums); !utils.SlicesEqualUnordered2D(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
