package main

import "testing"

type testCase struct {
	str1   string
	str2   string
	expect bool
}

func TestCanMakeSubsequence(t *testing.T) {
	testCases := []testCase{
		{"abc", "ad", true},
		{"zc", "ad", true},
		{"ab", "d", false},
		{"ff", "fg", true},
		{"eao", "ofa", false},
	}

	for _, tc := range testCases {
		if got := canMakeSubsequence(tc.str1, tc.str2); got != tc.expect {
			t.Errorf("canMakeSubsequence(%q, %q): got %v, expect %v", tc.str1, tc.str2, got, tc.expect)
		}
	}
}
