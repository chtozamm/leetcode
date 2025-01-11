package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	root     *TreeNode
	expected int
}

var testCases = []testCase{
	{root: treeFromArray([]int{3, 9, 20, -1, -1, 15, 7}, 0), expected: 3},
	{root: treeFromArray([]int{1, -1, 2}, 0), expected: 2},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := maxDepth(tc.root)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
