// 209. Minimum Size Subarray Sum
package min_subarr_len

func minSubArrayLen(target int, nums []int) int {
	minLen := len(nums) + 1
	leftIndex := 0
	sum := 0

	for rightIndex, num := range nums {
		sum += num

		for sum >= target {
			minLen = min(minLen, rightIndex-leftIndex+1)
			sum -= nums[leftIndex]
			leftIndex++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
