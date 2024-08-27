package justification

import (
	"reflect"
	"testing"
)

type testCase struct {
	words    []string
	maxWidth int
	expected []string
}

func TestFullJustify(t *testing.T) {
	testCases := []testCase{
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			expected: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			expected: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			expected: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for _, tc := range testCases {
		if got := fullJustify(tc.words, tc.maxWidth); !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
