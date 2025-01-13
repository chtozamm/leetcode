package main

func search(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n

	for low < high {
		mid := (low + high) / 2

		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low > 0 && nums[low-1] == target {
		return low - 1
	}

	return -1
}
