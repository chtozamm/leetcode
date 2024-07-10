package utils

import "math"

// WithinTolerance checks if two floating-point numbers are within a specified tolerance.
// It returns true if the absolute difference between the numbers is less than the tolerance,
// or if the relative difference is within the tolerance for non-zero values of b.
func WithinTolerance(a, b, e float64) bool {
	// Handle exact equality
	if a == b {
		return true
	}

	// Calculate the absolute difference
	d := math.Abs(a - b)

	// Handle the case where b is zero
	if b == 0 {
		return d < e
	}

	// Compare the relative difference
	return d/math.Abs(b) < e
}

// SlicesEqualUnordered checks if two slices are equal ignoring the order.
func SlicesEqualUnordered(slice1, slice2 []int) bool {
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
