package valid_parenthesis

import "testing"

type testCase struct {
	s        string
	expected bool
}

func TestIsValid(t *testing.T) {
	testCases := []testCase{
		{"()", true},
		{"(]", false},
		{"([)]", false},
		{"{[()]}", true},
		{"{[()]})", false},
	}

	for _, tc := range testCases {
		if got := isValid(tc.s); got != tc.expected {
			t.Errorf("isValid(%q) = %v, want %v", tc.s, got, tc.expected)
		}
	}
}
