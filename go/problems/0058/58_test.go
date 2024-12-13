package last_word

import "testing"

type testCase struct {
	s        string
	expected int
}

func TestLengthOfLastWord(t *testing.T) {
	testCases := []testCase{
		{s: "Hello World", expected: 5},
		{s: "   fly me   to   the moon  ", expected: 4},
		{s: "luffy is still joyboy", expected: 6},
	}

	for _, tc := range testCases {
		if got := lengthOfLastWord(tc.s); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
