package main

func waysToSplitArray(nums []int) int {
	n := len(nums)
	ans := 0

	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0

	for i := 0; i < n-1; i++ {
		leftSum += nums[i]
		rightSum := totalSum - leftSum

		if leftSum >= rightSum {
			ans++
		}
	}

	return ans
}
