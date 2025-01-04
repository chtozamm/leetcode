package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s        string
	expected int
}

var testCases = []testCase{
	{s: "aabca", expected: 3},
	{s: "adc", expected: 0},
	{s: "bbcbaba", expected: 4},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := countPalindromicSubsequence(tc.s)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
