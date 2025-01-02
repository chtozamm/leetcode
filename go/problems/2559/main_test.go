package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	words    []string
	queries  [][]int
	expected []int
}

var testCases = []testCase{
	{
		words:    []string{"aba", "bcb", "ece", "aa", "e"},
		queries:  [][]int{{0, 2}, {1, 4}, {1, 1}},
		expected: []int{2, 3, 0},
	},
	{
		words:    []string{"a", "e", "i"},
		queries:  [][]int{{0, 2}, {0, 1}, {2, 2}},
		expected: []int{3, 2, 1},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := vowelStrings(tc.words, tc.queries)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
