package jump_game_ii

import "testing"

type testCase struct {
	nums     []int
	expected int
}

func TestCanJump(t *testing.T) {
	testCases := []testCase{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{3, 2, 1, 0, 4}, 0},
	}

	for _, tc := range testCases {
		got := jump(tc.nums)
		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
