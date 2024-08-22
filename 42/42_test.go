package rain_water

import "testing"

type testCase struct {
	height   []int
	expected int
}

func TestTrap(t *testing.T) {
	testCases := []testCase{
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, expected: 6},
		{height: []int{4, 2, 0, 3, 2, 5}, expected: 9},
	}

	for _, tc := range testCases {
		if got := trap(tc.height); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
