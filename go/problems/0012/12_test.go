package int_to_roman

import "testing"

type testCase struct {
	num      int
	expected string
}

func TestIntToRoman(t *testing.T) {
	testCases := []testCase{
		{num: -19, expected: ""},
		{num: 0, expected: ""},
		{num: 1, expected: "I"},
		{num: 3, expected: "III"},
		{num: 4, expected: "IV"},
		{num: 7, expected: "VII"},
		{num: 9, expected: "IX"},
		{num: 58, expected: "LVIII"},
		{num: 1994, expected: "MCMXCIV"},
		{num: 3744, expected: "MMMDCCXLIV"},
	}

	for _, tc := range testCases {
		got := intToRoman(tc.num)
		if got != tc.expected {
			t.Errorf("intToRoman(%d) = %s; expected %s", tc.num, got, tc.expected)
		}
	}
}
