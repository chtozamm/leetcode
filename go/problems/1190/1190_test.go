package parentheses

import "testing"

type testCase struct {
	input    string
	expected string
}

func TestReverseParentheses(t *testing.T) {
	testCases := []testCase{
		{input: "(abcd)", expected: "dcba"},
		{input: "(u(love)i)", expected: "iloveu"},
		{input: "(ed(et(oc))el)", expected: "leetcode"},
	}

	for _, tc := range testCases {
		got := reverseParentheses(tc.input)
		if got != tc.expected {
			t.Errorf("reverseParentheses(%q) = %q; expected %q", tc.input, got, tc.expected)
		}
	}
}
