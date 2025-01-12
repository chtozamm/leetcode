package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	str1     string
	str2     string
	expected string
}

var testCases = []testCase{
	{str1: "ABCABC", str2: "ABC", expected: "ABC"},
	{str1: "ABABAB", str2: "ABAB", expected: "AB"},
	{str1: "LEET", str2: "CODE", expected: ""},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := gcdOfStrings(tc.str1, tc.str2)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
