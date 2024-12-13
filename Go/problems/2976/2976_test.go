package minimum_cost

import "testing"

type testCase struct {
	source   string
	target   string
	original []byte
	changed  []byte
	cost     []int
	expected int64
}

func TestMinimumCost(t *testing.T) {
	testCases := []testCase{
		{source: "abcd", target: "acbe", original: []byte{'a', 'b', 'c', 'c', 'e', 'd'}, changed: []byte{'b', 'c', 'b', 'e', 'b', 'e'}, cost: []int{2, 5, 5, 1, 2, 20}, expected: 28},
		{source: "aaaa", target: "bbbb", original: []byte{'a', 'c'}, changed: []byte{'c', 'b'}, cost: []int{1, 2}, expected: 12},
		{source: "abcd", target: "abce", original: []byte{'a'}, changed: []byte{'e'}, cost: []int{10000}, expected: -1},
	}

	for _, tc := range testCases {
		got := minimumCost(tc.source, tc.target, tc.original, tc.changed, tc.cost)

		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
