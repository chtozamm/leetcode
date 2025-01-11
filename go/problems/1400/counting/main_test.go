package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s        string
	k        int
	expected bool
}

var testCases = []testCase{
	{s: "anna", k: 1, expected: true},
	{s: "annabelle", k: 2, expected: true},
	{s: "leetcode", k: 3, expected: false},
	{s: "true", k: 4, expected: true},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := canConstruct(tc.s, tc.k)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
