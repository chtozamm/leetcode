package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	grid     [][]int
	expected int64
}

var testCases = []testCase{
	{
		grid:     [][]int{{2, 5, 4}, {1, 5, 1}},
		expected: 4,
	},
	{
		grid:     [][]int{{3, 3, 1}, {8, 5, 2}},
		expected: 4,
	},
	{
		grid:     [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}},
		expected: 7,
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := gridGame(tc.grid)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
