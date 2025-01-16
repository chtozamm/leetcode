package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums1    []int
	nums2    []int
	expected int
}

var testCases = []testCase{
	{
		nums1:    []int{2, 1, 3},
		nums2:    []int{10, 2, 5, 0},
		expected: 13,
	},
	{
		nums1:    []int{1, 2},
		nums2:    []int{3, 4},
		expected: 0,
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := xorAllNums(tc.nums1, tc.nums2)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
