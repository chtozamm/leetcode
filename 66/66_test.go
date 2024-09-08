package plus_one

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		digits   []int
		expected []int
	}{
		{digits: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{digits: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}},
		{digits: []int{9}, expected: []int{1, 0}},
		{digits: []int{1, 9, 9}, expected: []int{2, 0, 0}},
	}

	for i, tc := range testCases {
		got := plusOne(tc.digits)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("Test case %d failed: got %v, expected %v", i+1, got, tc.expected)
		}
	}
}
