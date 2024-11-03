package candy

import "testing"

type testCase struct {
	ratings  []int
	expected int
}

func TestCandy(t *testing.T) {
	testCases := []testCase{
		{ratings: []int{1, 0, 2}, expected: 5},
		{ratings: []int{1, 2, 2}, expected: 4},
	}

	for _, tc := range testCases {
		got := candy(tc.ratings)
		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
