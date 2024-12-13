package people_heights

import (
	"reflect"
	"testing"
)

type testCase struct {
	names    []string
	heights  []int
	expected []string
}

func TestSortPeople(t *testing.T) {
	testCases := []testCase{
		{names: []string{"Mary", "John", "Emma"}, heights: []int{180, 165, 170}, expected: []string{"Mary", "Emma", "John"}},
		{names: []string{"Alice", "Bob", "Bob"}, heights: []int{155, 185, 150}, expected: []string{"Bob", "Alice", "Bob"}},
	}

	for _, tc := range testCases {
		got := sortPeople(tc.names, tc.heights)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}
