package rotate_string

import "testing"

type testCase struct {
	s        string
	goal     string
	expected bool
}

func TestRotateString(t *testing.T) {
	testCases := []testCase{
		{s: "abcde", goal: "abcde", expected: true},
		{s: "abcde", goal: "abced", expected: false},
		{s: "defdefdefabcabc", goal: "defdefabcabcdef", expected: true},
	}

	for _, tc := range testCases {
		if got := rotateString(tc.s, tc.goal); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
