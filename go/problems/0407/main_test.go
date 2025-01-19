package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	heightMap [][]int
	expected  int
}

var testCases = []testCase{
	{
		heightMap: [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}},
		expected:  4,
	},
	{
		heightMap: [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}},
		expected:  10,
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := trapRainWater(tc.heightMap)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
