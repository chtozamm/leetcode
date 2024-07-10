package crawler

import (
	"testing"
)

type testCase struct {
	input    []string
	expected int
}

func TestMinOperations(t *testing.T) {
	testCases := []testCase{
		{input: []string{"d1/", "d2/", "../", "d21/", "./"}, expected: 2},
		{input: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}, expected: 3},
		{input: []string{"d1/", "../", "../", "../"}, expected: 0},
	}

	for _, tc := range testCases {
		got := minOperations(tc.input)
		if got != tc.expected {
			t.Errorf("got %d, expected %d", got, tc.expected)
		}
	}
}
