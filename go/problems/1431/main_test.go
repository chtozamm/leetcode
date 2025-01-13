package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	candies      []int
	extraCandies int
	expected     []bool
}

var testCases = []testCase{
	{
		candies:      []int{2, 3, 5, 1, 3},
		extraCandies: 3,
		expected:     []bool{true, true, true, false, true},
	},
	{
		candies:      []int{4, 2, 1, 1, 2},
		extraCandies: 1,
		expected:     []bool{true, false, false, false, false},
	},
	{
		candies:      []int{12, 1, 12},
		extraCandies: 10,
		expected:     []bool{true, false, true},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := kidsWithCandies(tc.candies, tc.extraCandies)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
