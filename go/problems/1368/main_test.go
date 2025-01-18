package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	grid     [][]int
	expected int
}

var testCases = []testCase{
	{grid: [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}, expected: 3},
	{grid: [][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}, expected: 0},
	{grid: [][]int{{1, 2}, {4, 3}}, expected: 1},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := minCost(tc.grid)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
