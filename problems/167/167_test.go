package two_sum

import (
	"reflect"
	"testing"
)

type testCase struct {
	numbers  []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	testCases := []testCase{
		{numbers: []int{2, 7, 11, 15}, target: 9, expected: []int{1, 2}},
		{numbers: []int{2, 3, 4}, target: 6, expected: []int{1, 3}},
		{numbers: []int{-1, 0}, target: -1, expected: []int{1, 2}},
	}

	for _, tc := range testCases {
		if got := twoSum(tc.numbers, tc.target); !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
