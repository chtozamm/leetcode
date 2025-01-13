package main

func search(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n

	for low < high {
		mid := (low + high) / 2

		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if low < n && nums[low] == target {
		return low
	}

	return -1
}
