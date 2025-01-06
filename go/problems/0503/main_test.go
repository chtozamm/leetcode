package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	expected []int
}

var testCases = []testCase{
	{nums: []int{1, 2, 1}, expected: []int{2, -1, 2}},
	{nums: []int{1, 2, 3, 4, 3}, expected: []int{2, 3, 4, -1, 4}},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := nextGreaterElements(tc.nums)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
