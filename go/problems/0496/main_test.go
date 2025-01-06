package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	nums1    []int
	nums2    []int
	expected []int
}

var testCases = []testCase{
	{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}, expected: []int{-1, 3, -1}},
	{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}, expected: []int{3, -1}},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := nextGreaterElement(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
