package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	words    []string
	pref     string
	expected int
}

var testCases = []testCase{
	{words: []string{"pay", "attention", "practice", "attend"}, pref: "at", expected: 2},
	{words: []string{"leetcode", "win", "loops", "success"}, pref: "code", expected: 0},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := prefixCount(tc.words, tc.pref)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
