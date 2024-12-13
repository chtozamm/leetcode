package add_binary

import "testing"

type testCase struct {
	a        string
	b        string
	expected string
}

func TestAddBinary(t *testing.T) {
	testCases := []testCase{
		{a: "11", b: "1", expected: "100"},
	}

	for _, tc := range testCases {
		got := addBinary(tc.a, tc.b)
		if got != tc.expected {
			t.Errorf("got %q, expected %q", got, tc.expected)
		}
	}
}
