package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	arr      []int
	mat      [][]int
	expected int
}

var testCases = []testCase{
	{
		arr:      []int{1, 3, 4, 2},
		mat:      [][]int{{1, 4}, {2, 3}},
		expected: 2,
	},
	{
		arr:      []int{2, 8, 7, 4, 1, 3, 5, 6, 9},
		mat:      [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}},
		expected: 3,
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := firstCompleteIndex(tc.arr, tc.mat)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
