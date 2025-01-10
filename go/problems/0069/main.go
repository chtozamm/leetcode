package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	low, high := 1, x

	for low < high {
		mid := low + (high-low)/2

		if mid*mid > x {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low - 1
}
