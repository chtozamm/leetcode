package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	boxes    string
	expected []int
}

var testCases = []testCase{
	{boxes: "110", expected: []int{1, 1, 3}},
	{boxes: "001011", expected: []int{11, 8, 5, 4, 3, 4}},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := minOperations(tc.boxes)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
