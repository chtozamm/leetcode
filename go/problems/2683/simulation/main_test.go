package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	derived  []int
	expected bool
}

var testCases = []testCase{
	{derived: []int{1, 1, 0}, expected: true},
	{derived: []int{1, 1}, expected: true},
	{derived: []int{1, 0}, expected: false},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := doesValidArrayExist(tc.derived)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
