package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	n        int
	expected int
}

var testCases = []testCase{
	{n: 1, expected: 1},
	{n: 2, expected: 2},
	{n: 3, expected: 3},
	{n: 4, expected: 5},
	{n: 5, expected: 8},
	{n: 6, expected: 13},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := climbStairs(tc.n)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
