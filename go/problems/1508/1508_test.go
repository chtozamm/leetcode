package range_sum

import "testing"

type testCase struct {
	nums     []int
	n        int
	left     int
	right    int
	expected int
}

func TestRangeSum(t *testing.T) {
	testCases := []testCase{
		{nums: []int{1, 2, 3, 4}, n: 4, left: 1, right: 5, expected: 13},
		{nums: []int{1, 2, 3, 4}, n: 4, left: 3, right: 4, expected: 6},
		{nums: []int{1, 2, 3, 4}, n: 4, left: 1, right: 10, expected: 50},
	}

	for _, tc := range testCases {
		got := rangeSum(tc.nums, tc.n, tc.left, tc.right)
		if got != tc.expected {
			t.Errorf("rangeSumBruteForce(%v, %d, %d, %d) = %d, expected %d", tc.nums, tc.n, tc.left, tc.right, got, tc.expected)
		}
	}
}
