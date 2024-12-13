package ransom_note

import (
	"testing"
)

type testCase struct {
	ransomNote string
	magazine   string
	expected   bool
}

func TestCanConstruct(t *testing.T) {
	testCases := []testCase{
		{ransomNote: "a", magazine: "b", expected: false},
		{ransomNote: "aa", magazine: "ab", expected: false},
		{ransomNote: "aa", magazine: "aab", expected: true},
	}

	for _, tc := range testCases {
		if got := canConstruct(tc.ransomNote, tc.magazine); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
