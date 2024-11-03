package pillow

import "testing"

type testCase struct {
	n        int
	time     int
	expected int
}

func TestPassThePillow(t *testing.T) {
	testCases := []testCase{
		{n: 4, time: 5, expected: 2},
		{n: 3, time: 2, expected: 3},
	}

	for _, tc := range testCases {
		got := PassThePillow(tc.n, tc.time)
		if got != tc.expected {
			t.Errorf("PassThePillow(%d, %d) = %d; expected %d", tc.n, tc.time, got, tc.expected)
		}
	}
}
