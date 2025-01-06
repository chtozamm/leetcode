package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	nums       []int
	k          int
	multiplier int
	expected   []int
}

var testCases = []testCase{
	{nums: []int{2, 1, 3, 5, 6}, k: 5, multiplier: 2, expected: []int{8, 4, 6, 5, 6}},
	{nums: []int{1, 2}, k: 3, multiplier: 4, expected: []int{16, 8}},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := getFinalState(tc.nums, tc.k, tc.multiplier)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
