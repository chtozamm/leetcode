package minimum_time

import "testing"

type testCase struct {
	n        int
	edges    [][]int
	time     int
	change   int
	expected int
}

func TestSecondMinimum(t *testing.T) {
	testCases := []testCase{
		{n: 5, edges: [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, time: 3, change: 5, expected: 13},
		{n: 2, edges: [][]int{{1, 2}}, time: 3, change: 2, expected: 11},
	}

	for _, tc := range testCases {
		t.Run("Breadth First Search", func(t *testing.T) {
			got := secondMinimumBFS(tc.n, tc.edges, tc.time, tc.change)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
		t.Run("Dijkstra", func(t *testing.T) {
			got := secondMinimumDijkstra(tc.n, tc.edges, tc.time, tc.change)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
