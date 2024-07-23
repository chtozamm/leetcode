package frequency_sort

import (
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	expected []int
}

func TestFrequencySort(t *testing.T) {
	testCases := []testCase{
		{nums: []int{1, 1, 2, 2, 2, 3}, expected: []int{3, 1, 1, 2, 2, 2}},
		{nums: []int{2, 3, 1, 3, 2}, expected: []int{1, 3, 3, 2, 2}},
		{nums: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, expected: []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
	}

	for _, tc := range testCases {
		got := frequencySort(tc.nums)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
