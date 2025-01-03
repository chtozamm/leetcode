package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums     []int
	expected int
}

var testCases = []testCase{
	{nums: []int{10, 4, -8, 7}, expected: 2},
	{nums: []int{2, 3, 1, 0}, expected: 2},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := waysToSplitArray(tc.nums)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
