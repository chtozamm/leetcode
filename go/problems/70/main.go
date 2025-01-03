package main

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}

	current := 0
	previous := 3
	twoStepsBack := 2

	for i := 4; i <= n; i++ {
		current = previous + twoStepsBack
		twoStepsBack = previous
		previous = current
	}

	return current
}
