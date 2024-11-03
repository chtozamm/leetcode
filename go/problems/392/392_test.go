package subsequence

import "testing"

type testCase struct {
	s        string
	t        string
	expected bool
}

func TestIsSubsequence(t *testing.T) {
	testCases := []testCase{
		{s: "abc", t: "ahbgdc", expected: true},
		{s: "axc", t: "ahbgdc", expected: false},
	}

	for _, tc := range testCases {
		if got := isSubsequence(tc.s, tc.t); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
