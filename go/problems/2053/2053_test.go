package distinct_string

import "testing"

type testCase struct {
	arr      []string
	k        int
	expected string
}

func TestKthDistinct(t *testing.T) {
	testCases := []testCase{
		{arr: []string{"d", "b", "c", "b", "c", "a"}, k: 2, expected: "a"},
		{arr: []string{"aaa", "aa", "a"}, k: 1, expected: "aaa"},
		{arr: []string{"a", "b", "a"}, k: 3, expected: ""},
	}

	for _, tc := range testCases {
		got := kthDistinct(tc.arr, tc.k)
		if got != tc.expected {
			t.Errorf("kthDistinct(%v, %d) = %s; expected %s", tc.arr, tc.k, got, tc.expected)
		}
	}
}
