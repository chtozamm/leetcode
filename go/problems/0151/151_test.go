package reverse_words

import "testing"

type testCase struct {
	s        string
	expected string
}

func TestReverseWords(t *testing.T) {
	testCases := []testCase{
		{s: "the sky is blue", expected: "blue is sky the"},
		{s: "  hello world  ", expected: "world hello"},
		{s: "a good   example", expected: "example good a"},
	}

	for _, tc := range testCases {
		if got := reverseWords(tc.s); got != tc.expected {
			t.Errorf("got %q, expected %q", got, tc.expected)
		}
	}
}
