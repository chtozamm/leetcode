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
	{grid: [][]int{{1, 0}, {0, 1}}, expected: 0},
	{grid: [][]int{{1, 0}, {1, 1}}, expected: 3},
	{grid: [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}, expected: 4},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := countServers(tc.grid)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
