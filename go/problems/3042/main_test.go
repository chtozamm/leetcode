package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	words    []string
	expected int
}

var testCases = []testCase{
	{words: []string{"a", "aba", "ababa", "aa"}, expected: 4},
	{words: []string{"pa", "papa", "ma", "mama"}, expected: 2},
	{words: []string{"abab", "ab"}, expected: 0},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := countPrefixSuffixPairs(tc.words)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
