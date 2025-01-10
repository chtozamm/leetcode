package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	words1   []string
	words2   []string
	expected []string
}

var testCases = []testCase{
	{
		words1:   []string{"amazon", "apple", "facebook", "google", "leetcode"},
		words2:   []string{"e", "o"},
		expected: []string{"facebook", "google", "leetcode"},
	},
	{
		words1:   []string{"amazon", "apple", "facebook", "google", "leetcode"},
		words2:   []string{"l", "e"},
		expected: []string{"apple", "google", "leetcode"},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := wordSubsets(tc.words1, tc.words2)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
