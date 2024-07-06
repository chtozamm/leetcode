package array

import (
	"testing"
)

type testCase struct {
	arr1     []int
	arr2     []int
	expected []int
}

func TestMinDifference(t *testing.T) {
	testCases := []testCase{
		{arr1: []int{1, 2, 2, 1}, arr2: []int{2, 2}, expected: []int{2, 2}},
		{arr1: []int{4, 9, 5}, arr2: []int{9, 4, 9, 8, 4}, expected: []int{4, 9}},
		{arr1: []int{3, 2, 3, 5}, arr2: []int{1, 6, 4, 3, 3, 2}, expected: []int{3, 3, 2}},
	}

	for _, tc := range testCases {
		got := intersect(tc.arr1, tc.arr2)
		if !slicesEqualUnordered(got, tc.expected) {
			t.Errorf("intersect(%v, %v) = %v; expected %v", tc.arr1, tc.arr2, got, tc.expected)
		}
	}
}

// slicesEqualUnordered checks if two slices are equal ignoring the order.
func slicesEqualUnordered(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	counts := make(map[int]int)

	for _, num := range slice1 {
		counts[num]++
	}

	for _, num := range slice2 {
		counts[num]--
		if counts[num] < 0 {
			return false
		}
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
