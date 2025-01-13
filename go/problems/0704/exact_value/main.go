package main

func search(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n-1

	for low <= high {
		mid := (low + high) / 2
		curr := nums[mid]

		if curr == target {
			return mid
		} else if curr > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
