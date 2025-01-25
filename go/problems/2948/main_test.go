package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	limit    int
	expected []int
}

var testCases = []testCase{
	{
		nums:     []int{1, 5, 3, 9, 8},
		limit:    2,
		expected: []int{1, 3, 5, 8, 9},
	},
	{
		nums:     []int{1, 7, 6, 18, 2, 1},
		limit:    3,
		expected: []int{1, 6, 7, 18, 1, 2},
	},
	{
		nums:     []int{1, 7, 28, 19, 1},
		limit:    3,
		expected: []int{1, 7, 28, 19, 1},
	},
	{
		nums:     []int{1, 81, 10, 79, 36, 2, 87, 12, 20, 77},
		limit:    7,
		expected: []int{1, 77, 10, 79, 36, 2, 81, 12, 20, 87},
	},
	{
		nums:     []int{4, 52, 38, 59, 71, 27, 31, 83, 88, 10},
		limit:    14,
		expected: []int{4, 27, 31, 38, 52, 59, 71, 83, 88, 10},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := lexicographicallySmallestArray(tc.nums, tc.limit)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
