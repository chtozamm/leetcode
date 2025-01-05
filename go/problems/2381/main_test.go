package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s        string
	shifts   [][]int
	expected string
}

var testCases = []testCase{
	{
		s:        "abc",
		shifts:   [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}},
		expected: "ace",
	},
	{
		s:        "dztz",
		shifts:   [][]int{{0, 0, 0}, {1, 1, 1}},
		expected: "catz",
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := shiftingLetters(tc.s, tc.shifts)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
