// 1509. Minimum Difference Between Largest and Smallest Value in Three Moves
package distance

import (
	"math"
	"slices"
)

func minDifference(nums []int) int {
	n := len(nums)
	if n <= 4 {
		return 0
	}

	slices.Sort(nums)
	result := math.Inf(1)

	for i := 0; i <= 3; i++ {
		result = math.Min(result, float64(nums[n-4+i]-nums[i]))
	}

	return int(result)
}
