package gain

import "testing"

type testCase struct {
	s        string
	x        int
	y        int
	expected int
}

func TestMaximumGain(t *testing.T) {
	testCases := []testCase{
		{s: "cdbcbbaaabab", x: 4, y: 5, expected: 19},
		{s: "aabbaaxybbaabb", x: 5, y: 4, expected: 20},
	}

	for _, tc := range testCases {
		got := maximumGain(tc.s, tc.x, tc.y)
		if got != tc.expected {
			t.Errorf("got %d, expected %d", got, tc.expected)
		}
	}
}
