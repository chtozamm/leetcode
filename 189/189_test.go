package rotate_array

import (
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	k        int
	expected []int
}

func TestRotate(t *testing.T) {
	testCases := []testCase{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, expected: []int{5, 6, 7, 1, 2, 3, 4}},
		{nums: []int{-1, -100, 3, 99}, k: 2, expected: []int{3, 99, -1, -100}},
	}

	for _, tc := range testCases {
		rotate(tc.nums, tc.k)
		if !reflect.DeepEqual(tc.nums, tc.expected) {
			t.Errorf("got %v, expected %v", tc.nums, tc.expected)
		}
	}
}
