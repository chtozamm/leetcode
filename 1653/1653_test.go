package balanced_string

import "testing"

type testCase struct {
	s        string
	expected int
}

func TestMinimumDeletions(t *testing.T) {
	testCases := []testCase{
		{"aababbab", 2},
		{"bbaaaaabb", 2},
	}

	for _, tc := range testCases {
		got := minimumDeletions(tc.s)
		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
