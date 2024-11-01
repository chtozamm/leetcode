package common_prefix

import "testing"

type testCase struct {
	strs     []string
	expected string
}

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []testCase{
		{strs: []string{"flower", "flow", "flight"}, expected: "fl"},
		{strs: []string{"dog", "racecar", "car"}, expected: ""},
	}

	for _, tc := range testCases {
		if got := longestCommonPrefix(tc.strs); got != tc.expected {
			t.Errorf("got %q, expected %q", got, tc.expected)
		}
	}
}
