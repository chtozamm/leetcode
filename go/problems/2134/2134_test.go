package min_swaps

import "testing"

type testCase struct {
	nums     []int
	expected int
}

func TestMinSwaps(t *testing.T) {
	testCases := []testCase{
		{nums: []int{0, 1, 0, 1, 1, 0, 0}, expected: 1},
		{nums: []int{0, 1, 1, 1, 0, 0, 1, 1, 0}, expected: 2},
		{nums: []int{1, 1, 0, 0, 1}, expected: 0},
	}

	for _, tc := range testCases {
		got := minSwaps(tc.nums)
		if got != tc.expected {
			t.Errorf("got %d, expected %d", got, tc.expected)
		}
	}
}
