package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	words    []string
	expected []string
}

var testCases = []testCase{
	{words: []string{"mass", "as", "hero", "superhero"}, expected: []string{"as", "hero"}},
	{words: []string{"leetcode", "et", "code"}, expected: []string{"et", "code"}},
	{words: []string{"blue", "green", "bu"}, expected: []string{}},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := stringMatching(tc.words)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
