package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s        string
	expected bool
}

var testCases = []testCase{
	{s: "aba", expected: true},
	{s: "abca", expected: true},
	{s: "abc", expected: false},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := validPalindrome(tc.s)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
