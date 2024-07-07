package water

import "testing"

type testCase struct {
	numBottles  int
	numExchange int
	expected    int
}

func TestNumWaterBottles(t *testing.T) {
	testCases := []testCase{
		{numBottles: 9, numExchange: 3, expected: 13},
		{numBottles: 15, numExchange: 4, expected: 19},
	}

	for _, tc := range testCases {
		got := numWaterBottles(tc.numBottles, tc.numExchange)
		if got != tc.expected {
			t.Errorf("numWaterBottles(%d, %d) = %d; expected %d", tc.numBottles, tc.numExchange, got, tc.expected)
		}
	}
}
