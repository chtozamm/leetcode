package min_subarr_len

import "testing"

type testCase struct {
	target   int
	nums     []int
	expected int
}

func TestMinSubArrayLen(t *testing.T) {
	testCases := []testCase{
		{target: 7, nums: []int{2, 3, 1, 2, 4, 3}, expected: 2},
		{target: 4, nums: []int{1, 4, 4}, expected: 1},
		{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}, expected: 0},
	}

	for _, tc := range testCases {
		got := minSubArrayLen(tc.target, tc.nums)
		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
