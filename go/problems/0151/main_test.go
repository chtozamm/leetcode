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
	{s: "the sky is blue", expected: "blue is sky the"},
	{s: "  hello world  ", expected: "world hello"},
	{s: "a good   example", expected: "example good a"},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := reverseWords(tc.s)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
