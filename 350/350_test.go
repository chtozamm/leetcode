package array

import (
	"testing"

	"github.com/chtozamm/leetcode/utils"
)

type testCase struct {
	arr1     []int
	arr2     []int
	expected []int
}

func TestIntersect(t *testing.T) {
	testCases := []testCase{
		{arr1: []int{1, 2, 2, 1}, arr2: []int{2, 2}, expected: []int{2, 2}},
		{arr1: []int{4, 9, 5}, arr2: []int{9, 4, 9, 8, 4}, expected: []int{4, 9}},
		{arr1: []int{3, 2, 3, 5}, arr2: []int{1, 6, 4, 3, 3, 2}, expected: []int{3, 3, 2}},
	}

	for _, tc := range testCases {
		got := intersect(tc.arr1, tc.arr2)
		if !utils.SlicesEqualUnordered(got, tc.expected) {
			t.Errorf("intersect(%v, %v) = %v; expected %v", tc.arr1, tc.arr2, got, tc.expected)
		}
	}
}
