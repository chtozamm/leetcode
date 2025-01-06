package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	gifts    []int
	k        int
	expected int64
}

var testCases = []testCase{
	{gifts: []int{25, 64, 9, 4, 100}, k: 4, expected: 29},
	{gifts: []int{1, 1, 1, 1}, k: 4, expected: 4},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := pickGifts(tc.gifts, tc.k)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
