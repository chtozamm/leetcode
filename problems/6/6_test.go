package zigzag_pattern

import "testing"

type testCase struct {
	s        string
	numRows  int
	expected string
}

func TestConvert(t *testing.T) {
	testCases := []testCase{
		{s: "PAYPALISHIRING", numRows: 3, expected: "PAHNAPLSIIGYIR"},
		{s: "PAYPALISHIRING", numRows: 4, expected: "PINALSIGYAHRPI"},
		{s: "A", numRows: 1, expected: "A"},
	}

	for _, tc := range testCases {
		if got := convert(tc.s, tc.numRows); got != tc.expected {
			t.Errorf("got %q, expected %q", got, tc.expected)
		}
	}
}
