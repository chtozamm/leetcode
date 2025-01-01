package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	s        string
	expected int
}

var testCases = []testCase{
	{"011101", 5},
	{"00111", 5},
	{"1111", 3},
	{"0000000000", 9},
	{"1111111111", 9},
	{"0101010101", 6},
	{"1010101010", 5},
	{"000111000111", 9},
	{"111000111000", 6},
	{"01010101010101010101", 11},
	{"11111111110000000000", 9},
	{"00000000001111111111", 20},
	{"0101010101010101010101010101010101", 18},
	{"0000000000000000000000000000000000000000", 39},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := maxScore(tc.s)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			maxScore(tc.s)
		}
	}
}
