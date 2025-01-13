package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s        string
	expected string
}

var testCases = []testCase{
	{s: "IceCreAm", expected: "AceCreIm"},
	{s: "leetcode", expected: "leotcede"},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := reverseVowels(tc.s)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
