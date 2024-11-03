package circular_game

import "testing"

type testCase struct {
	n        int
	k        int
	expected int
}

func TestFindTheWinner(t *testing.T) {
	testCases := []testCase{
		{n: 5, k: 2, expected: 3},
		{n: 6, k: 5, expected: 1},
		{n: 10, k: 1, expected: 10},
	}

	for _, tc := range testCases {
		got := findTheWinner(tc.n, tc.k)
		if got != tc.expected {
			t.Errorf("findTheWinner(%d, %d) = %d; expected %d", tc.n, tc.k, got, tc.expected)
		}
	}
}
