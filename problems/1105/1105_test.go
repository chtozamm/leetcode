package bookshelf

import "testing"

type testCase struct {
	books      [][]int
	shelfWidth int
	expected   int
}

func TestMinHeightShelves(t *testing.T) {
	testCases := []testCase{
		{books: [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, shelfWidth: 4, expected: 6},
		{books: [][]int{{1, 3}, {2, 4}, {3, 2}}, shelfWidth: 6, expected: 4},
	}

	for _, tc := range testCases {
		got := minHeightShelves(tc.books, tc.shelfWidth)

		if got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
