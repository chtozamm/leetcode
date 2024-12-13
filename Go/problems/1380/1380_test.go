package lucky_numbers

import (
	"reflect"
	"testing"
)

type testCase struct {
	input    [][]int
	expected []int
}

func TestLuckyNumbers(t *testing.T) {
	testCases := []testCase{
		{input: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, expected: []int{15}},
		{input: [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, expected: []int{12}},
		{input: [][]int{{7, 8}, {1, 2}}, expected: []int{7}},
	}

	for _, tc := range testCases {
		got := luckyNumbers(tc.input)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
