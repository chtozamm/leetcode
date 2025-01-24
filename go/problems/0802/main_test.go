package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	graph    [][]int
	expected []int
}

var testCases = []testCase{
	{
		graph:    [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}},
		expected: []int{2, 4, 5, 6},
	},
	{
		graph:    [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}},
		expected: []int{4},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := eventualSafeNodes(tc.graph)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
