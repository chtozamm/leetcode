package substring

import (
	"reflect"
	"testing"
)

type testCase struct {
	s        string
	words    []string
	expected []int
}

func TestFindSubstring(t *testing.T) {
	testCases := []testCase{
		{s: "barfoothefoobarman", words: []string{"foo", "bar"}, expected: []int{0, 9}},
		{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "word"}, expected: []int{}},
		{s: "barfoofoobarthefoobarman", words: []string{"bar", "foo", "the"}, expected: []int{6, 9, 12}},
	}

	for _, tc := range testCases {
		if got := findSubstring(tc.s, tc.words); !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %d, expected %d", got, tc.expected)
		}
	}
}
