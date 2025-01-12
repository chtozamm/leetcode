package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s        string
	locked   string
	expected bool
}

var testCases = []testCase{
	{s: "))()))", locked: "010100", expected: true},
	{s: "()()", locked: "0000", expected: true},
	{s: ")", locked: "0", expected: false},
	{s: ")(", locked: "00", expected: true},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := canBeValid(tc.s, tc.locked)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
