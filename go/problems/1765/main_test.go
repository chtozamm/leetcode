package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	isWater  [][]int
	expected [][]int
}

var testCases = []testCase{
	{
		isWater:  [][]int{{0, 1}, {0, 0}},
		expected: [][]int{{1, 0}, {2, 1}},
	},
	{
		isWater:  [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}},
		expected: [][]int{{1, 1, 0}, {0, 1, 1}, {1, 2, 2}},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := highestPeak(tc.isWater)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
