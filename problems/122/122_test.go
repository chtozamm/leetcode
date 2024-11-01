package profit

import "testing"

type testCase struct {
	prices   []int
	expected int
}

func TestMaxProfit(t *testing.T) {
	testCases := []testCase{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tc := range testCases {
		got := maxProfit(tc.prices)
		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
