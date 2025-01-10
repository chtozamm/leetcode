package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	x        float64
	n        int
	expected float64
}

var testCases = []testCase{
	{x: 2.00000, n: 10, expected: 1024.00000},
	{x: 2.10000, n: 3, expected: 9.26100},
	{x: 2.00000, n: -2, expected: 0.25000},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := myPow(tc.x, tc.n)
			if got > tc.expected+0.00001 || got < tc.expected-0.00001 {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
