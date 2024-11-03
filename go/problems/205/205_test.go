package isomophic_strings

import (
	"testing"
)

type testCase struct {
	s        string
	t        string
	expected bool
}

func TestIsIsomorphic(t *testing.T) {
	testCases := []testCase{
		{s: "egg", t: "add", expected: true},
		{s: "foo", t: "bar", expected: false},
		{s: "paper", t: "title", expected: true},
	}

	for _, tc := range testCases {
		if got := isIsomorphic(tc.s, tc.t); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
