package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	temperatures []int
	expected     []int
}

var testCases = []testCase{
	{temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73}, expected: []int{1, 1, 4, 2, 1, 1, 0, 0}},
	{temperatures: []int{30, 40, 50, 60}, expected: []int{1, 1, 1, 0}},
	{temperatures: []int{30, 60, 90}, expected: []int{1, 1, 0}},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := dailyTemperatures(tc.temperatures)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
