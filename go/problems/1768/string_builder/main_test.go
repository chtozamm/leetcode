package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	word1    string
	word2    string
	expected string
}

var testCases = []testCase{
	{word1: "abc", word2: "pqr", expected: "apbqcr"},
	{word1: "ab", word2: "pqrs", expected: "apbqrs"},
	{word1: "abcd", word2: "pq", expected: "apbqcd"},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := mergeAlternately(tc.word1, tc.word2)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
