package restaurant

import (
	"testing"

	"github.com/chtozamm/leetcode/utils"
)

type testCase struct {
	customers [][]int
	expected  float64
}

func TestFindTheWinner(t *testing.T) {
	testCases := []testCase{
		{customers: [][]int{{1, 2}, {2, 5}, {4, 3}}, expected: 5.0},
		{customers: [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}, expected: 3.25},
		{customers: [][]int{{2, 3}, {6, 3}, {7, 5}, {11, 3}, {15, 2}, {18, 1}}, expected: 4.16667},
	}

	for _, tc := range testCases {
		got := averageWaitingTime(tc.customers)
		if !utils.WithinTolerance(got, tc.expected, 5) {
			t.Errorf("got %.5f; expected %.5f", got, tc.expected)
		}
	}
}
