package palindrome

import "testing"

type testCase struct {
	x        int
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	testCases := []testCase{
		{x: 121, expected: true},
		{x: -121, expected: false},
		{x: 10, expected: false},
		{x: 5, expected: true},
		{x: 898, expected: true},
		{x: 555, expected: true},
		{x: 782, expected: false},
	}

	for _, tc := range testCases {
		got := isPalindrome(tc.x)
		if got != tc.expected {
			t.Errorf("isPalindrome(%d) = %v; expected %v", tc.x, got, tc.expected)
		}
	}
}
