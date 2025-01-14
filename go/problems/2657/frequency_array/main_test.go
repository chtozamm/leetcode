package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	A        []int
	B        []int
	expected []int
}

var testCases = []testCase{
	{A: []int{1, 3, 2, 4}, B: []int{3, 1, 2, 4}, expected: []int{0, 2, 3, 4}},
	{A: []int{2, 3, 1}, B: []int{3, 1, 2}, expected: []int{0, 1, 3}},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := findThePrefixCommonArray(tc.A, tc.B)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
