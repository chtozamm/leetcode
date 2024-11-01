package jump_game

import "testing"

type testCase struct {
	nums     []int
	expected bool
}

func TestCanJump(t *testing.T) {
	testCases := []testCase{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}

	for _, tc := range testCases {
		got := canJump(tc.nums)
		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
