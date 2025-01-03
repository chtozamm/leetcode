package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	low      int
	high     int
	zero     int
	one      int
	expected int
}

var testCases = []testCase{
	{low: 3, high: 3, zero: 1, one: 1, expected: 8},
	{low: 2, high: 3, zero: 1, one: 2, expected: 5},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := countGoodStrings(tc.low, tc.high, tc.zero, tc.one)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
