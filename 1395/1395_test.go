package count_teams

import "testing"

type testCase struct {
	rating   []int
	expected int
}

func TestNumTeams(t *testing.T) {
	testCases := []testCase{
		{rating: []int{2, 5, 3, 4, 1}, expected: 3},
		{rating: []int{2, 1, 3}, expected: 0},
		{rating: []int{1, 2, 3, 4}, expected: 4},
	}

	for _, tc := range testCases {
		got := numTeams(tc.rating)
		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
