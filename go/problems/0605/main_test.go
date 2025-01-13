package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	flowerbed []int
	n         int
	expected  bool
}

var testCases = []testCase{
	{flowerbed: []int{1, 0, 0, 0, 1}, n: 1, expected: true},
	{flowerbed: []int{1, 0, 0, 0, 1}, n: 2, expected: false},
	{flowerbed: []int{0}, n: 1, expected: true},
	{flowerbed: []int{1}, n: 1, expected: false},
	{flowerbed: []int{1}, n: 0, expected: true},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := canPlaceFlowers(tc.flowerbed, tc.n)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
