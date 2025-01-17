package main

// derived[i] = original[i] ^ original[i+1]
// original[i+1] = derived[i] ^ original[i]
// derived[n−1] = original[n−1] ^ original[0]

func doesValidArrayExist(derived []int) bool {
	n := len(derived)

	// Assume that the first element of the original array is 0
	original := make([]int, n+1)
	for i := 0; i < n; i++ {
		original[i+1] = derived[i] ^ original[i]
	}

	if original[0] == original[n] {
		return true
	}

	// Assume that the first element of the original array is 1
	original = make([]int, n+1)
	original[0] = 1
	for i := 0; i < n; i++ {
		original[i+1] = derived[i] ^ original[i]
	}

	if original[0] == original[n] {
		return true
	}

	return false
}
