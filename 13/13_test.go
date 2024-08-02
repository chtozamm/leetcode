package roman_to_int

import "testing"

type testCase struct {
	s        string
	expected int
}

func TestRomanToInt(t *testing.T) {
	testCases := []testCase{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"IV", 4},
		{"IX", 9},
	}

	for _, tc := range testCases {
		got := romanToInt(tc.s)
		if got != tc.expected {
			t.Errorf("RomanToInt(%s) = %d; expected %d", tc.s, got, tc.expected)
		}
	}

}
