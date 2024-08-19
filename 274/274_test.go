package h_index

import "testing"

type testCase struct {
	citations []int
	expected  int
}

func TestHIndex(t *testing.T) {
	testCases := []testCase{
		{citations: []int{3, 0, 6, 1, 5}, expected: 3},
		{citations: []int{1, 3, 1}, expected: 1},
		{citations: []int{0}, expected: 0},
	}

	for _, tc := range testCases {
		if got := hIndex(tc.citations); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
