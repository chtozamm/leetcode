// 2134. Minimum Swaps to Group All 1's Together II
package min_swaps

import "math"

func minSwaps(nums []int) int {
	minimumSwaps := math.MaxInt

	// Calculate the total number of 1s in the array
	totalOnes := 0
	for _, n := range nums {
		totalOnes += n
	}

	if totalOnes == 0 {
		return 0
	}

	n := len(nums)
	// Initialize the window of size totalOnes
	onesCount := 0
	for i := 0; i < totalOnes; i++ {
		onesCount += nums[i]
	}

	// Use a sliding window to find the minimum number of swaps
	for i := 0; i < n; i++ {
		minimumSwaps = min(minimumSwaps, totalOnes-onesCount)
		// Slide the window to the right by one position
		onesCount -= nums[i]
		onesCount += nums[(i+totalOnes)%n]
	}

	return minimumSwaps
}
