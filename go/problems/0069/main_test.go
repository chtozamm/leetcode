package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	x        int
	expected int
}

var testCases = []testCase{
	{x: 0, expected: 0},
	{x: 1, expected: 1},
	{x: 2, expected: 1},
	{x: 4, expected: 2},
	{x: 8, expected: 2},
	{x: 9, expected: 3},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := mySqrt(tc.x)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
