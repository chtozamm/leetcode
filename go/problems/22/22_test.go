package parentheses

import (
	"reflect"
	"testing"
)

type testCase struct {
	n        int
	expected []string
}

func TestGenerateParenthesis(t *testing.T) {
	testCases := []testCase{
		{n: 3, expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{n: 1, expected: []string{"()"}},
	}

	for _, tc := range testCases {
		got := generateParenthesis(tc.n)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("generateParenthesis(%d) = %v; expected %v", tc.n, got, tc.expected)
		}
	}
}
