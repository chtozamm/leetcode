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
	{s: "abaacbcbb", expected: 5},
	{s: "aa", expected: 2},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := minimumLength(tc.s)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
