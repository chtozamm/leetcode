package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	num1     int
	num2     int
	expected int
}

var testCases = []testCase{
	{num1: 3, num2: 5, expected: 3},
	{num1: 1, num2: 12, expected: 3},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := minimizeXor(tc.num1, tc.num2)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
