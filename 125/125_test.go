package palindrome

import "testing"

type testCase struct {
	s        string
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	testCases := []testCase{
		{s: "A man, a plan, a canal: Panama", expected: true},
		{s: "race a car", expected: false},
		{s: " ", expected: true},
	}

	for _, tc := range testCases {
		if got := isPalindrome(tc.s); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
