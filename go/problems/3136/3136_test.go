package valid_word

import "testing"

type testCase struct {
	input    string
	expected bool
}

func TestIsValid(t *testing.T) {
	testCases := []testCase{
		{input: "234Adas", expected: true},
		{input: "b3", expected: false},
		{input: "a3$e", expected: false},
	}

	for _, tc := range testCases {
		got := isValid(tc.input)
		if got != tc.expected {
			t.Errorf("isValid(%q) = %v; expected %v", tc.input, got, tc.expected)
		}
	}
}
