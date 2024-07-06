package distance

import "testing"

type testCase struct {
	nums     []int
	expected int
}

func TestMinDifference(t *testing.T) {
	testCases := []testCase{
		{nums: []int{5, 3, 2, 4}, expected: 0},
		{nums: []int{1, 5, 0, 10, 14}, expected: 1},
		{nums: []int{3, 100, 20}, expected: 0},
		{nums: []int{6, 6, 0, 1, 1, 4, 6}, expected: 2},
	}

	for _, tc := range testCases {
		got := minDifference(tc.nums)
		if got != tc.expected {
			t.Errorf("minDifference(%v) = %d; expected %d", tc.nums, got, tc.expected)
		}
	}
}
