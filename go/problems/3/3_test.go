package longest_substring

import "testing"

type testCase struct {
	s        string
	expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []testCase{
		{s: "abcabcbb", expected: 3},
		{s: "bbbbb", expected: 1},
		{s: "pwwkew", expected: 3},
	}

	for _, tc := range testCases {
		if got := lengthOfLongestSubstring(tc.s); got != tc.expected {
			t.Errorf("got %d, expected %d", got, tc.expected)
		}
	}
}
