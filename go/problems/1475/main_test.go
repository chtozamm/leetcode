package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	prices   []int
	expected []int
}

var testCases = []testCase{
	{prices: []int{8, 4, 6, 2, 3}, expected: []int{4, 2, 4, 2, 3}},
	{prices: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
	{prices: []int{10, 1, 1, 6}, expected: []int{9, 0, 1, 6}},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := finalPrices(tc.prices)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
