package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums     []int
	target   int
	expected int
}

var testCases = []testCase{
	{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, expected: 4},
	{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, expected: -1},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := search(tc.nums, tc.target)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
