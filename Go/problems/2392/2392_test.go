package matrix

import (
	"reflect"
	"testing"
)

type testCase struct {
	k             int
	rowConditions [][]int
	colConditions [][]int
	expected      [][]int
}

func TestBuildMatrix(t *testing.T) {
	testCases := []testCase{
		{k: 3, rowConditions: [][]int{{1, 2}, {3, 2}}, colConditions: [][]int{{2, 1}, {3, 2}}, expected: [][]int{{3, 0, 0}, {0, 0, 1}, {0, 2, 0}}},
		{k: 3, rowConditions: [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}, colConditions: [][]int{{2, 1}}, expected: [][]int{}},
	}
	for _, tc := range testCases {
		got := buildMatrix(tc.k, tc.rowConditions, tc.colConditions)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
