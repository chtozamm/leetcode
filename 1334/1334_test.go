package city_neighbors

import (
	"testing"
)

type testCase struct {
	n                 int
	edges             [][]int
	distanceThreshold int
	expected          int
}

func TestFindTheCity(t *testing.T) {
	testCases := []testCase{
		{n: 4, edges: [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, distanceThreshold: 4, expected: 3},
		{n: 5, edges: [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, distanceThreshold: 2, expected: 0},
	}

	for _, tc := range testCases {
		got := findTheCity(tc.n, tc.edges, tc.distanceThreshold)

		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
