package product_of_array

import (
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	expected []int
}

func TestProductExceptSelf(t *testing.T) {
	testCases := []testCase{
		{nums: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{nums: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
	}

	for _, tc := range testCases {
		if got := productExceptSelf(tc.nums); !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
