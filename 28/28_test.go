package occurence

import "testing"

type testCase struct {
	haystack string
	needle   string
	expected int
}

func TestStrStr(t *testing.T) {
	testCases := []testCase{
		{haystack: "sadbutsad", needle: "sad", expected: 0},
		{haystack: "leetcode", needle: "leeto", expected: -1},
	}

	for _, tc := range testCases {
		if got := strStr(tc.haystack, tc.needle); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
