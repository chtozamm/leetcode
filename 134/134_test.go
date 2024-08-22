package gas_station

import "testing"

type testCase struct {
	gas      []int
	cost     []int
	expected int
}

func TestCanCompleteCircuit(t *testing.T) {
	testCases := []testCase{
		{gas: []int{1, 2, 3, 4, 5}, cost: []int{3, 4, 5, 1, 2}, expected: 3},
		{gas: []int{2, 3, 4}, cost: []int{3, 4, 3}, expected: -1},
	}

	for _, tc := range testCases {
		if got := canCompleteCircuit(tc.gas, tc.cost); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
