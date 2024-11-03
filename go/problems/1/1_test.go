package two_sum

import (
	"testing"

	"github.com/chtozamm/leetcode/utils"
)

type testCase struct {
	nums     []int
	target   int
	expected []int
}

func TestTwoSums(t *testing.T) {
	testCases := []testCase{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		got := twoSum(tc.nums, tc.target)

		if !utils.SlicesEqualUnordered(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
