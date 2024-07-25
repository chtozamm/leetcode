package sort_array

import (
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	expected []int
}

func TestSortArray(t *testing.T) {
	testCases := []testCase{
		{nums: []int{5, 2, 3, 1}, expected: []int{1, 2, 3, 5}},
		{nums: []int{5, 1, 1, 2, 0, 0}, expected: []int{0, 0, 1, 1, 2, 5}},
		{nums: []int{4, -1, -3, 2, -2, 0, 4}, expected: []int{-3, -2, -1, 0, 2, 4, 4}},
	}

	for _, tc := range testCases {
		got := sortArray(tc.nums)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
