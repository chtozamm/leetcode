package seniors

import "testing"

type testCase struct {
	details  []string
	expected int
}

func TestCountSeniors(t *testing.T) {
	testCases := []testCase{
		{details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}, expected: 2},
	}

	for _, tc := range testCases {
		got := countSeniors(tc.details)
		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
