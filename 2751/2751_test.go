package robot_collisions

import (
	"reflect"
	"testing"
)

type testCase struct {
	positions  []int
	healths    []int
	directions string
	expected   []int
}

func TestSurvivedRobotsHealths(t *testing.T) {
	testCases := []testCase{
		{positions: []int{5, 4, 3, 2, 1}, healths: []int{2, 17, 9, 15, 10}, directions: "RRRRR", expected: []int{2, 17, 9, 15, 10}},
		{positions: []int{3, 5, 2, 6}, healths: []int{10, 10, 15, 12}, directions: "RLRL", expected: []int{14}},
		{positions: []int{1, 2, 5, 6}, healths: []int{10, 10, 11, 11}, directions: "RLRL", expected: []int{}},
		{positions: []int{11, 44, 16}, healths: []int{1, 20, 17}, directions: "RLR", expected: []int{18}},
	}

	for _, tc := range testCases {
		got := survivedRobotsHealths(tc.positions, tc.healths, tc.directions)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v; expected %v", got, tc.expected)
		}
	}
}
